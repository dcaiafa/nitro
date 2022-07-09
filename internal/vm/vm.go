package vm

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"github.com/dcaiafa/nitro/internal/fiber"
)

const stackSize = 1000

type Op byte

const (
	OpUMinus Op = iota
	OpAdd
	OpSub
	OpMult
	OpDiv
	OpMod
	OpLT
	OpLE
	OpGT
	OpGE
	OpEq
	OpNE
)

type FrameCrumb struct {
	fn    *Fn
	ip    int
	extFn Callable
}

func (c *FrameCrumb) IsZero() bool {
	return c.fn == nil && c.extFn == nil
}

type FrameInfo struct {
	Filename string
	Line     int
	Func     string
}

func (i FrameInfo) String() string {
	var loc string
	if i.Filename != "" {
		loc = fmt.Sprintf("%v:%v", i.Filename, i.Line)
	} else {
		loc = "<builtin>"
	}
	return fmt.Sprintf("%v  %v", loc, i.Func)
}

type tryCatch struct {
	CatchAddr int
}

type frame struct {
	nRet       int
	nArg       int
	nLocals    int
	iter       *ILIterator
	fn         *Fn
	extFn      Callable
	caps       []ValueRef
	tryCatches []tryCatch
	defers     []*Closure
	pipeline   bool
	ip         int
	bp         int
}

type VM struct {
	userData     map[interface{}]interface{}
	program      *CompiledPackage
	globals      []Value
	shuttingDown bool
	sched        *fiber.Scheduler
	co           *coroutine
	closers      map[Closer]struct{}
	errLogger    func(err error)
	firstErr     error

	mu          sync.Mutex
	interrupt   int32
	injectedErr error
}

func NewVM(prog *CompiledPackage) *VM {
	return &VM{
		program:   prog,
		globals:   make([]Value, prog.globals),
		userData:  make(map[interface{}]interface{}),
		sched:     fiber.NewScheduler(),
		closers:   make(map[Closer]struct{}),
		errLogger: func(err error) { fmt.Fprintln(os.Stderr, err) },
	}
}

func (m *VM) SetUserData(key, value interface{}) {
	m.userData[key] = value
}

func (m *VM) GetUserData(key interface{}) interface{} {
	return m.userData[key]
}

func (m *VM) SetParam(n string, v Value) error {
	param := m.program.params[n]
	if param == nil {
		return fmt.Errorf("invalid parameter: %s", n)
	}
	m.globals[param.global] = v
	return nil
}

func (m *VM) Run(args []Value) error {
	co := m.newCoroutine()
	co.PushContext(context.Background())
	copy(co.stack, args)
	co.sp = len(args)

	f := co.NewFrame()
	f.fn = &m.program.fns[m.program.MainFnNdx]
	f.nArg = len(args)
	f.bp = len(args)

	co.PushFrame(f)
	m.sched.Run(co.fiber)
	m.shutdown()

	return m.firstErr
}

func (m *VM) newCoroutine() *coroutine {
	co := &coroutine{}
	co.fiber = m.sched.NewFiber(func() {
		m.co = co
		err := m.resume()
		if err != nil {
			m.errLogger(err)
		}
		if m.firstErr == nil {
			m.firstErr = err
		}
	})
	co.fiber.Data = co
	co.stack = co.preAllocStack[:]
	return co
}

func (m *VM) StartCoroutine(callable Value) error {
	co := m.newCoroutine()
	co.PushContext(m.co.TopContext())

	f := co.NewFrame()

	switch callable := callable.(type) {
	case *Closure:
		f.fn = callable.fn
		f.caps = callable.caps

	case *Fn:
		f.fn = callable

	default:
		return fmt.Errorf("cannot start coroutine with %v", TypeName(callable))
	}

	co.PushFrame(f)
	prev := m.co
	m.sched.SwitchToNew(co.fiber)
	m.co = prev

	return nil
}

func (m *VM) Context() context.Context {
	return m.co.TopContext()
}

func (m *VM) PushContext(ctx context.Context) {
	m.co.PushContext(ctx)
}

func (m *VM) PopContext() context.Context {
	return m.co.PopContext()
}

func (m *VM) Block(f func(ctx context.Context)) {
	self := m.co

	m.sched.Block(self.TopContext(), func(ctx context.Context) {
		f(ctx)
	})
	m.co = self
}

func (m *VM) Call(callable Value, args []Value, nret int) ([]Value, error) {
	sp := m.co.sp
	copy(m.co.stack[m.co.sp:], args)
	m.co.sp += len(args)
	err := m.call(callable, len(args), nret, false)
	if err != nil {
		return nil, err
	}
	var ret []Value
	if nret > 0 {
		ret = make([]Value, nret)
	}
	copy(ret, m.co.stack[m.co.sp-nret:])
	m.co.sp = sp
	return ret, nil
}

func (m *VM) RegisterCloser(c Closer) {
	m.closers[c] = struct{}{}
}

func (m *VM) UnregisterCloser(c Closer) {
	delete(m.closers, c)
}

func (m *VM) callExtFn(
	extFn Callable,
	caps []ValueRef,
	narg int,
	nret int,
	pipeline bool,
) (err error) {
	f := m.co.NewFrame()
	f.nRet = nret
	f.nArg = narg
	f.extFn = extFn
	f.caps = caps
	f.bp = m.co.sp
	f.pipeline = pipeline

	m.co.PushFrame(f)

	args := m.co.stack[m.co.sp-narg : m.co.sp]

	rets, err := extFn.Call(m, args, nret)
	if err != nil {
		wrapRuntimeError(m, &err)
		m.co.PopFrame()
		return err
	}

	if len(rets) < nret {
		err = fmt.Errorf(
			"external function was expected to return at least %v values, "+
				"but it returned only %v values", nret, len(rets))
		wrapRuntimeError(m, &err)
		m.co.PopFrame()
		return err
	}

	if nret > 0 {
		copy(m.co.stack[m.co.sp-narg:], rets[:nret])
	}

	m.co.sp = m.co.sp - narg + nret

	m.co.PopFrame()
	return nil
}

func (m *VM) call(callable Value, narg int, nret int, pipeline bool) error {
	switch callable := callable.(type) {
	case *Closure:
		f := m.co.NewFrame()
		f.fn = callable.fn
		f.caps = callable.caps
		f.nArg = narg
		f.nRet = nret
		f.pipeline = pipeline
		return m.runFrame(f)

	case *NativeIterator:
		return m.callExtFn(callable.extFn, nil, narg, nret, false)

	case *ILIterator:
		if callable.ip == -1 {
			m.co.stack[m.co.sp] = False
			for i := 1; i < nret; i++ {
				m.co.stack[m.co.sp+i] = nil
			}
			m.co.sp += nret
			return nil
		}

		rsp := m.co.sp
		rstack := m.co.stack

		f := m.co.NewFrame()
		f.fn = callable.fn
		f.iter = callable
		f.caps = callable.captures
		f.tryCatches = callable.tryCatches
		f.defers = callable.defers
		f.nRet = nret
		f.nLocals = callable.nlocals
		f.ip = callable.ip

		m.co.stack = callable.stack
		m.co.sp = callable.sp

		err := m.runFrame(f)
		if err != nil {
			m.co.stack = rstack
			m.co.sp = rsp
			return err
		}

		copy(rstack[rsp:], m.co.stack[m.co.sp-nret:m.co.sp])
		callable.sp = m.co.sp - nret
		m.co.stack = rstack
		m.co.sp = rsp + nret
		return nil

	case *Fn:
		f := m.co.NewFrame()
		f.fn = callable
		f.nArg = narg
		f.nRet = nret
		f.pipeline = true
		return m.runFrame(f)

	case Callable:
		return m.callExtFn(callable, nil, narg, nret, pipeline)

	default:
		if callable == nil {
			return ErrCannotCallNil
		}
		return fmt.Errorf("cannot call %q", TypeName(callable))
	}
}

func (m *VM) IterNext(iter Iterator, nret int) ([]Value, error) {
	if nret == 0 {
		return nil, errors.New("nret is zero")
	}

	sp := m.co.sp
	ok, err := m.iterNext(iter, nret)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	ret := make([]Value, nret)
	copy(ret, m.co.stack[m.co.sp-nret:])
	m.co.sp = sp
	return ret, nil
}

func (m *VM) IterClose(iter Iterator) error {
	nativeIter, ok := iter.(*NativeIterator)
	if !ok {
		return nil
	}

	return nativeIter.Close(m)
}

func (m *VM) ShuttingDown() bool {
	return m.shuttingDown
}

func (m *VM) iterNext(iter Iterator, nret int) (bool, error) {
	switch iter := iter.(type) {
	case *NativeIterator:
		f := m.co.NewFrame()
		f.nRet = nret
		f.extFn = iter.extFn
		f.bp = m.co.sp
		m.co.PushFrame(f)

		rets, err := iter.extFn.Call(m, nil, nret)
		if err != nil {
			wrapRuntimeError(m, &err)
			m.co.PopFrame()
			iter.Close(m)
			return false, err
		}

		if len(rets) == 0 {
			m.co.PopFrame()
			return false, iter.Close(m)
		}

		if len(rets) < nret {
			err = fmt.Errorf(
				"native iterator was expected to return at least %v values, "+
					"but it returned only %v values", nret, len(rets))
			iter.Close(m)
			wrapRuntimeError(m, &err)
			m.co.PopFrame()
			return false, err
		}

		copy(m.co.stack[m.co.sp:], rets[:nret])
		m.co.sp = m.co.sp + nret
		m.co.PopFrame()
		return true, nil

	case *ILIterator:
		if iter.ip == -1 {
			return false, nil
		}

		rsp := m.co.sp
		rstack := m.co.stack

		f := m.co.NewFrame()
		f.fn = iter.fn
		f.iter = iter
		f.caps = iter.captures
		f.tryCatches = iter.tryCatches
		f.defers = iter.defers
		f.nRet = nret
		f.nLocals = iter.nlocals
		f.ip = iter.ip

		m.co.stack = iter.stack
		m.co.sp = iter.sp

		err := m.runFrame(f)
		if err != nil {
			m.co.stack = rstack
			m.co.sp = rsp
			return false, err
		}

		if iter.ip == -1 {
			m.co.stack = rstack
			m.co.sp = rsp
			return false, nil
		}

		copy(rstack[rsp:], m.co.stack[m.co.sp-nret:m.co.sp])
		iter.sp = m.co.sp - nret
		m.co.stack = rstack
		m.co.sp = rsp + nret
		return true, nil

	default:
		if iter == nil {
			return false, ErrCannotCallNil
		}
		return false, fmt.Errorf("cannot call %q", TypeName(iter))
	}
}

func (m *VM) runFrame(frame *frame) error {
	m.co.PushFrame(frame)
	return m.resume()
}

func (m *VM) resume() (err error) {
	for {
		err := m.resumeWithoutRecovery()
		if err == nil {
			err = m.runDefers()
			m.co.PopFrame()
			return err
		}

		// Update the current frame's ip so that the stack trace created by
		// wrapRuntimeError will reflect the current position.
		m.co.frame.ip = m.co.ip
		rerr := wrapRuntimeError(m, &err)
		err = rerr

		if len(m.co.frame.tryCatches) == 0 {
			derr := m.runDefers()
			if derr != nil {
				err = fmt.Errorf(
					"defer threw error:\n%v\n"+
						"while handling error:\n%w",
					derr, err)
			}
			m.co.PopFrame()
			return err
		}

		tryCatch := m.co.frame.tryCatches[len(m.co.frame.tryCatches)-1]
		m.co.frame.tryCatches = m.co.frame.tryCatches[:len(m.co.frame.tryCatches)-1]
		m.co.ip = tryCatch.CatchAddr
		m.co.sp = m.co.frame.bp + m.co.frame.nLocals + 1
		m.co.stack[m.co.sp-1] = rerr
	}
}

func (m *VM) runDefers() error {
	for i := len(m.co.frame.defers) - 1; i >= 0; i-- {
		deferred := m.co.frame.defers[i]
		_, err := m.Call(deferred, nil, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *VM) processInterrupt() {
	if atomic.LoadInt32(&m.interrupt) == 0 {
		return
	}

	atomic.StoreInt32(&m.interrupt, 0)

	var err error
	m.mu.Lock()
	if m.injectedErr != nil {
		err = m.injectedErr
		m.injectedErr = nil
	}
	m.mu.Unlock()

	if err != nil {
		m.sched.ForEachFiber(func(f *fiber.Fiber) {
			co := f.Data.(*coroutine)
			co.pendingErr = err
		})
	}
}

func (m *VM) resumeWithoutRecovery() (err error) {
	defer func() {
		m.processInterrupt()
		if m.co.pendingErr != nil {
			err = m.co.pendingErr
			m.co.pendingErr = nil
		}
	}()

	for {
		m.processInterrupt()
		if m.co.pendingErr != nil {
			return m.co.pendingErr
		}

		instr := m.co.instrs[m.co.ip]

		switch instr.opc {
		case OpNop:

		case OpJump:
			m.co.ip = int(instr.op1) - 1

		case OpJumpIfTrue:
			v := m.co.stack[m.co.sp-1]
			m.co.sp--
			b := CoerceToBool(v)
			if b {
				m.co.ip = int(instr.op1) - 1
			}

		case OpJumpIfFalse:
			v := m.co.stack[m.co.sp-1]
			m.co.sp--
			b := CoerceToBool(v)
			if !b {
				m.co.ip = int(instr.op1) - 1
			}

		case OpDup:
			m.co.stack[m.co.sp] = m.co.stack[m.co.sp-1]
			m.co.sp++

		case OpPop:
			m.co.sp -= int(instr.op1)

		case OpCall:
			nret := int(instr.op2)
			narg := int(instr.op1 & CallArgCountMask)
			expand := (instr.op1 & CallExpandFlag) != 0
			pipeline := (instr.op1 & CallPipelineFlag) != 0

			if expand {
				if narg == 0 {
					return fmt.Errorf("assert: zero arguments in expansion")
				}
				lastArg := m.co.stack[m.co.sp-1]
				if lastArg != nil {
					arr, ok := lastArg.(*Array)
					if !ok {
						return fmt.Errorf("cannot expand %v argument", TypeName(lastArg))
					}
					copy(m.co.stack[m.co.sp-1:], arr.array)
					m.co.sp += len(arr.array) - 1
					narg += len(arr.array) - 1
				} else {
					m.co.sp--
					narg--
				}
			}

			callable := m.co.stack[m.co.sp-narg-1]
			rsp := m.co.sp - narg - 1 + nret
			err = m.call(callable, narg, nret, pipeline)
			if err != nil {
				return err
			}
			if nret > 0 {
				rets := m.co.stack[m.co.sp-nret : m.co.sp]
				copy(m.co.stack[rsp-nret:], rets)
			}
			m.co.sp = rsp

		case OpNil:
			m.co.stack[m.co.sp] = nil
			m.co.sp++

		case OpNewClosure:
			capN := int(instr.op2)
			fn := int(instr.op1)

			var caps []ValueRef
			if capN > 0 {
				caps = make([]ValueRef, capN)
				for i, capture := range m.co.stack[m.co.sp-capN : m.co.sp] {
					caps[i] = capture.(ValueRef)
				}
				m.co.sp -= capN
			}

			closure := &Closure{
				fn:   &m.program.fns[fn],
				caps: caps,
			}
			m.co.stack[m.co.sp] = closure
			m.co.sp++

		case OpNewIter:
			capN := int(instr.op2)
			fn := int(instr.op1 & 0x00ffffff)
			iterNRet := int(instr.op1 & 0xff000000 >> 24)

			var caps []ValueRef
			if capN > 0 {
				caps = make([]ValueRef, capN)
				for i, capture := range m.co.stack[m.co.sp-capN : m.co.sp] {
					caps[i] = capture.(ValueRef)
				}
				m.co.sp -= capN
			}

			iter := &ILIterator{
				fn:       &m.program.fns[fn],
				captures: caps,
				iterNRet: iterNRet,
			}
			iter.stack = iter.preAllocStack[:]

			m.co.stack[m.co.sp] = iter
			m.co.sp++

		case OpNewInt:
			m.co.stack[m.co.sp] = NewInt(int64(instr.op1))
			m.co.sp++

		case OpNewBool:
			m.co.stack[m.co.sp] = NewBool(instr.op1 != 0)
			m.co.sp++

		case OpNewObject:
			m.co.stack[m.co.sp] = NewObject()
			m.co.sp++

		case OpNewArray:
			m.co.stack[m.co.sp] = NewArray()
			m.co.sp++

		case OpLoadGlobal:
			m.co.stack[m.co.sp] = m.globals[int(instr.op1)]
			m.co.sp++

		case OpLoadGlobalRef:
			m.co.stack[m.co.sp] = ValueRef{&m.globals[int(instr.op1)]}
			m.co.sp++

		case OpLoadGlobalDeref:
			m.co.stack[m.co.sp] = *m.globals[int(instr.op1)].(ValueRef).Ref
			m.co.sp++

		case OpLoadLocal:
			m.co.stack[m.co.sp] = m.co.stack[m.co.frame.bp+int(instr.op1)]
			m.co.sp++

		case OpLoadLocalRef:
			m.co.stack[m.co.sp] = ValueRef{&m.co.stack[m.co.frame.bp+int(instr.op1)]}
			m.co.sp++

		case OpLoadLocalDeref:
			m.co.stack[m.co.sp] = *m.co.stack[m.co.frame.bp+int(instr.op1)].(ValueRef).Ref
			m.co.sp++

		case OpCaptureLocal:
			l := m.co.stack[m.co.frame.bp+int(instr.op1)]
			if _, ok := l.(ValueRef); !ok {
				ref := ValueRef{Ref: new(Value)}
				*ref.Ref = l
				m.co.stack[m.co.frame.bp+int(instr.op1)] = ref
				l = ref
			}
			m.co.stack[m.co.sp] = l
			m.co.sp++

		case OpLoadArg:
			idx := int(instr.op1)
			if idx < m.co.frame.nArg {
				m.co.stack[m.co.sp] = m.co.stack[m.co.frame.bp-m.co.frame.nArg+idx]
			} else {
				m.co.stack[m.co.sp] = nil
			}
			m.co.sp++

		case OpLoadArgRef:
			idx := int(instr.op1)
			if idx >= m.co.frame.nArg {
				// TODO: min arg count
				return fmt.Errorf(
					"cannot assign to arg because it was not provided by caller")
			}
			m.co.stack[m.co.sp] = ValueRef{&m.co.stack[m.co.frame.bp-m.co.frame.nArg+idx]}
			m.co.sp++

		case OpLoadArgDeref:
			idx := int(instr.op1)
			if idx < m.co.frame.nArg {
				m.co.stack[m.co.sp] = *m.co.stack[m.co.frame.bp-m.co.frame.nArg+idx].(ValueRef).Ref
			} else {
				m.co.stack[m.co.sp] = nil
			}
			m.co.sp++

		case OpCaptureArg:
			a := m.co.stack[m.co.frame.bp-m.co.frame.nArg+int(instr.op1)]
			if _, ok := a.(ValueRef); !ok {
				ref := ValueRef{Ref: new(Value)}
				*ref.Ref = a
				m.co.stack[m.co.frame.bp-m.co.frame.nArg+int(instr.op1)] = ref
				a = ref
			}
			m.co.stack[m.co.sp] = a
			m.co.sp++

		case OpLoadCapture:
			m.co.stack[m.co.sp] = *m.co.frame.caps[int(instr.op1)].Ref
			m.co.sp++

		case OpLoadCaptureRef:
			m.co.stack[m.co.sp] = m.co.frame.caps[int(instr.op1)]
			m.co.sp++

		case OpLoadFn:
			m.co.stack[m.co.sp] = &m.program.fns[int(instr.op1)]
			m.co.sp++

		case OpLoadNativeFn:
			m.co.stack[m.co.sp] = m.program.extFns[int(instr.op1)]
			m.co.sp++

		case OpLoadLiteral:
			m.co.stack[m.co.sp] = m.program.literals[int(instr.op1)]
			m.co.sp++

		case OpEvalBinOp:
			op := Op(instr.op1)
			operand1 := m.co.stack[m.co.sp-2]
			operand2 := m.co.stack[m.co.sp-1]
			res, err := EvalOp(op, operand1, operand2)
			if err != nil {
				return err
			}
			m.co.stack[m.co.sp-2] = res
			m.co.sp--

		case OpNot:
			term := m.co.stack[m.co.sp-1]
			m.co.stack[m.co.sp-1] = NewBool(!CoerceToBool(term))

		case OpUnaryMinus:
			term := m.co.stack[m.co.sp-1]
			if term == nil {
				return errors.New("value is nil")
			}
			res, err := EvalOp(OpUMinus, term, nil)
			if err != nil {
				return err
			}
			m.co.stack[m.co.sp-1] = res

		case OpObjectPutNoPop:
			obj := m.co.stack[m.co.sp-3].(*Object)
			key := m.co.stack[m.co.sp-2]
			val := m.co.stack[m.co.sp-1]
			obj.Put(key, val)
			m.co.sp -= 2

		case OpObjectGet:
			objRaw := m.co.stack[m.co.sp-2]
			key := m.co.stack[m.co.sp-1]
			if objRaw == nil {
				m.co.stack[m.co.sp-2] = nil
			} else {
				indexable, ok := objRaw.(Indexable)
				if !ok {
					return fmt.Errorf("type %v is not indexable", TypeName(objRaw))
				}
				value, err := indexable.Index(key)
				if err != nil {
					return err
				}
				m.co.stack[m.co.sp-2] = value
			}
			m.co.sp--

		case OpObjectGetRef:
			objRaw := m.co.stack[m.co.sp-2]
			key := m.co.stack[m.co.sp-1]
			if objRaw == nil {
				return fmt.Errorf("cannot assign: value is nil")
			}
			indexable, ok := objRaw.(Indexable)
			if !ok {
				return fmt.Errorf("type %v is not indexable", TypeName(objRaw))
			}
			valueRef, err := indexable.IndexRef(key)
			if err != nil {
				return err
			}
			m.co.stack[m.co.sp-2] = valueRef
			m.co.sp--

		case OpArrayAppendNoPop:
			array := m.co.stack[m.co.sp-2].(*Array)
			value := m.co.stack[m.co.sp-1]
			array.Add(value)
			m.co.sp--

		case OpRet:
			nret := m.co.sp - (m.co.frame.bp + m.co.frame.nLocals)
			if m.co.frame.nRet > nret {
				return fmt.Errorf(
					"function was expected to return at least %v values, "+
						"but returned only %v",
					m.co.frame.nRet, nret)
			}
			if nret > m.co.frame.nRet {
				m.co.sp -= (nret - m.co.frame.nRet)
			}
			return nil

		case OpStore:
			count := int(instr.op1)
			for i := 0; i < count; i++ {
				rval := m.co.stack[m.co.sp-(count*2-i)].(ValueRef)
				val := m.co.stack[m.co.sp-(count-i)]
				*rval.Ref = val
			}
			m.co.sp -= count * 2

		case OpInitCallFrame:
			m.co.frame.nLocals = int(instr.op1)
			m.co.frame.bp = m.co.sp
			m.co.sp += m.co.frame.nLocals
			for i := m.co.frame.bp; i < m.co.sp; i++ {
				m.co.stack[i] = nil
			}

		case OpMakeIter:
			v := m.co.stack[m.co.sp-1]
			switch v := v.(type) {
			case Iterator:
				// Ready to go.
			case Iterable:
				m.co.stack[m.co.sp-1] = v.MakeIterator()
			default:
				if v == nil {
					m.co.stack[m.co.sp-1] = NewIterator(emptyIter, nil, 1)
				} else {
					return fmt.Errorf(
						"cannot iterate over value of type %q", TypeName(v))
				}
			}

		case OpBeginTry:
			m.co.frame.tryCatches = append(m.co.frame.tryCatches, tryCatch{
				CatchAddr: int(instr.op1),
			})

		case OpEndTry:
			m.co.frame.tryCatches = m.co.frame.tryCatches[:len(m.co.frame.tryCatches)-1]
			m.co.ip = int(instr.op1) - 1

		case OpSwap:
			i1 := m.co.sp - 1
			i2 := i1 - int(instr.op1)
			m.co.stack[i1], m.co.stack[i2] = m.co.stack[i2], m.co.stack[i1]

		case OpThrow:
			errVal := m.co.stack[m.co.sp-1]
			m.co.sp--
			err, ok := errVal.(*RuntimeError)
			if !ok {
				err = &RuntimeError{ErrValue: errVal}
			}
			return err

		case OpDefer:
			deferClosure := m.co.stack[m.co.sp-1].(*Closure)
			m.co.sp--
			m.co.frame.defers = append(m.co.frame.defers, deferClosure)

		case OpNext:
			iter, ok := m.co.stack[m.co.sp-1].(Iterator)
			if !ok {
				return fmt.Errorf("%q is not an iterator", TypeName(m.co.stack[m.co.sp-1]))
			}

			jumpTo := int(instr.op1)
			n := int(instr.op2)

			rsp := m.co.sp - n - 1

			// With n = 3:
			// rsp  +0  +1  +2  +3  +4  +5  +6
			//      r1  r2  r3  it               before iterNext
			//      r1  r2  r3  v1  v2  v3       after iterNext

			ok, err := m.iterNext(iter, n)
			if err != nil {
				return err
			}

			if ok {
				for i := 0; i < n; i++ {
					rval := m.co.stack[rsp+i].(ValueRef)
					val := m.co.stack[rsp+n+1+i]
					*rval.Ref = val
				}
			} else {
				m.co.ip = jumpTo - 1
			}

			m.co.sp = rsp

		case OpSlice:
			target := m.co.stack[m.co.sp-3]
			begin := m.co.stack[m.co.sp-2]
			end := m.co.stack[m.co.sp-1]

			sliceable, ok := target.(interface {
				Slice(begin Value, end Value) (Value, error)
			})
			if !ok {
				return fmt.Errorf("cannot slice %q", TypeName(target))
			}

			res, err := sliceable.Slice(begin, end)
			if err != nil {
				return err
			}
			m.co.stack[m.co.sp-3] = res
			m.co.sp -= 2

		case OpIterYield:
			nret := m.co.sp - (m.co.frame.bp + m.co.frame.nLocals)
			if nret < m.co.frame.nRet {
				return fmt.Errorf(
					"iterator was expected to yield at least %v values, "+
						"but yielded only %v",
					m.co.frame.nRet, nret)
			}

			iter := m.co.frame.iter
			iter.tryCatches = m.co.frame.tryCatches
			iter.defers = m.co.frame.defers
			iter.nlocals = m.co.frame.nLocals
			iter.ip = m.co.ip + 1

			if nret > m.co.frame.nRet {
				m.co.sp -= (nret - m.co.frame.nRet)
			}

			return nil

		case OpIterRet:
			m.co.frame.iter.ip = -1
			return nil

		case OpLiftArg:
			idx := int(instr.op1)
			if idx < m.co.frame.nArg {
				lifted := ValueRef{new(Value)}
				*lifted.Ref = m.co.stack[m.co.frame.bp-m.co.frame.nArg+idx]
				m.co.stack[m.co.frame.bp-m.co.frame.nArg+idx] = lifted
			}

		case OpInitLocal:
			m.co.stack[m.co.frame.bp+int(instr.op1)] = nil

		case OpInitLiftedLocal:
			m.co.stack[m.co.frame.bp+int(instr.op1)] = ValueRef{new(Value)}

		case OpInitGlobal:
			m.globals[int(instr.op1)] = nil

		case OpInitLiftedGlobal:
			m.globals[int(instr.op1)] = ValueRef{new(Value)}

		default:
			panic("invalid instruction")
		}

		m.co.ip++
	}
}

func (m *VM) SignalError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.injectedErr = err

	atomic.StoreInt32(&m.interrupt, 1)

	m.sched.CancelBlocked()
}

func (m *VM) GetCallerNArg() int {
	if len(m.co.callStack) < 2 {
		return 0
	}
	return m.co.callStack[len(m.co.callStack)-2].nArg
}

func (m *VM) GetCallerArgs() []Value {
	if len(m.co.callStack) < 2 {
		return nil
	}
	f := m.co.callStack[len(m.co.callStack)-2]
	args := m.co.stack[f.bp-f.nArg : f.bp]
	return args
}

func (m *VM) IsPipeline() bool {
	if len(m.co.callStack) < 1 {
		return false
	}
	return m.co.callStack[len(m.co.callStack)-1].pipeline
}

func (m *VM) IsCallerPipeline() bool {
	if len(m.co.callStack) < 2 {
		return false
	}
	return m.co.callStack[len(m.co.callStack)-2].pipeline
}

func (m *VM) GetStackInfo() []FrameInfo {
	stack := make([]FrameInfo, 0, len(m.co.callStack))
	for i := 0; i < len(m.co.callStack); i++ {
		stack = append(stack, m.GetFrameInfo(m.GetFrameCrumb(i)))
	}
	return stack
}

func (m *VM) GetFrameCrumb(n int) FrameCrumb {
	if n >= len(m.co.callStack) {
		return FrameCrumb{}
	}
	frame := m.co.callStack[len(m.co.callStack)-n-1]
	return FrameCrumb{
		fn:    frame.fn,
		ip:    frame.ip,
		extFn: frame.extFn,
	}
}

func (m *VM) GetFrameInfo(crumb FrameCrumb) FrameInfo {
	if crumb.fn != nil {
		loc := m.getLocation(crumb.fn, crumb.ip)
		if loc == nil {
			return FrameInfo{
				Filename: "???",
				Line:     0,
				Func:     "???",
			}
		}
		return FrameInfo{
			Filename: m.program.literals[loc.filename].(String).String(),
			Line:     loc.lineNum,
			Func:     m.program.literals[loc.fn].(String).String(),
		}
	}
	if nativeFn, ok := crumb.extFn.(*NativeFn); ok {
		return FrameInfo{
			Func: nativeFn.Name(),
		}
	}
	return FrameInfo{
		Filename: "???",
		Line:     0,
		Func:     "???",
	}
}

func (m *VM) getLocation(fn *Fn, ip int) *Location {
	locs := fn.locations
	if len(fn.locations) == 0 {
		return nil
	}
	for i := range locs {
		if locs[i].ip > ip {
			if i == 0 {
				return nil
			}
			return &locs[i-1]
		}
	}

	return &locs[len(fn.locations)-1]
}

func (m *VM) shutdown() {
	m.shuttingDown = true
	for c := range m.closers {
		c.Close()
	}
}
