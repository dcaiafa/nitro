package vm

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const stackSize = 1000

type BinOp byte

const (
	BinAdd BinOp = iota
	BinSub
	BinMult
	BinDiv
	BinMod
	BinLT
	BinLE
	BinGT
	BinGE
	BinEq
	BinNE
)

type FrameInfo struct {
	Filename string
	Line     int
	Func     string
}

type tryCatch struct {
	CatchAddr int
}

type frame struct {
	nRet       int
	nArg       int
	nLocals    int
	iter       *Iterator
	fn         *Fn
	extFn      NativeFn
	caps       []ValueRef
	tryCatches []tryCatch
	defers     []*Closure
	ip         int
	bp         int
}

type VM struct {
	ctx       context.Context
	userData  map[interface{}]interface{}
	program   *Program
	globals   []Value
	callStack []*frame
	frame     *frame
	stack     []Value
	instrs    []Instr
	sp        int
	ip        int
	framePool []*frame

	preAllocStack [stackSize]Value
}

func NewVM(ctx context.Context, prog *Program) *VM {
	m := &VM{
		ctx:      ctx,
		program:  prog,
		globals:  make([]Value, prog.globals),
		userData: make(map[interface{}]interface{}),
	}
	m.stack = m.preAllocStack[:]
	return m
}

func (m *VM) Context() context.Context {
	return m.ctx
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

func (m *VM) Run() error {
	f := m.newFrame()
	f.fn = &m.program.fns[0]
	return m.runFrame(f)
}

func (m *VM) Call(callable Value, args []Value, nret int) ([]Value, error) {
	sp := m.sp
	copy(m.stack[m.sp:], args)
	m.sp += len(args)
	err := m.call(callable, len(args), nret)
	if err != nil {
		return nil, err
	}
	var ret []Value
	if nret > 0 {
		ret = make([]Value, nret)
	}
	copy(ret, m.stack[m.sp-nret:])
	m.sp = sp
	return ret, nil
}

func (m *VM) callExtFn(
	extFn NativeFn,
	caps []ValueRef,
	narg int,
	nret int,
) (err error) {
	f := m.newFrame()
	f.nRet = nret
	f.nArg = narg
	f.extFn = extFn
	f.caps = caps
	f.bp = m.sp

	m.pushFrame(f)
	defer m.popFrame()
	defer wrapRuntimeError(m, &err)

	args := m.stack[m.sp-narg : m.sp]

	rets, err := extFn(m, args, nret)
	if err != nil {
		return err
	}
	if len(rets) < nret {
		return fmt.Errorf(
			"external function was expected to return at least %v values, "+
				"but it returned only %v values", nret, len(rets))
	}

	if nret > 0 {
		copy(m.stack[m.sp-narg:], rets[:nret])
	}
	m.sp = m.sp - narg + nret

	return nil
}

func (m *VM) call(callable Value, narg int, nret int) error {
	switch callable := callable.(type) {
	case *Closure:
		if callable.extFn != nil {
			return m.callExtFn(callable.extFn, callable.caps, narg, nret)
		} else {
			f := m.newFrame()
			f.fn = callable.fn
			f.caps = callable.caps
			f.nArg = narg
			f.nRet = nret
			return m.runFrame(f)
		}

	case *Iterator:
		if callable.extFn != nil {
			return m.callExtFn(callable.extFn, callable.captures, narg, nret)
		} else {
			if callable.ip == -1 {
				m.stack[m.sp] = False
				for i := 1; i < nret; i++ {
					m.stack[m.sp+i] = nil
				}
				m.sp += nret
				return nil
			}

			rsp := m.sp
			rstack := m.stack

			f := m.newFrame()
			f.fn = callable.fn
			f.iter = callable
			f.caps = callable.captures
			f.tryCatches = callable.tryCatches
			f.defers = callable.defers
			f.nRet = nret
			f.nLocals = callable.nlocals
			f.ip = callable.ip

			m.stack = callable.stack
			m.sp = callable.sp

			err := m.runFrame(f)
			if err != nil {
				m.stack = rstack
				m.sp = rsp
				return err
			}

			copy(rstack[rsp:], m.stack[m.sp-nret:m.sp])
			callable.sp = m.sp - nret
			m.stack = rstack
			m.sp = rsp + nret
			return nil
		}

	case *Fn:
		f := m.newFrame()
		f.fn = callable
		f.nArg = narg
		f.nRet = nret
		return m.runFrame(f)

	case NativeFn:
		return m.callExtFn(callable, nil, narg, nret)

	default:
		if callable == nil {
			return ErrCannotCallNil
		}
		return fmt.Errorf("cannot call type %q", TypeName(callable))
	}
}

func (m *VM) pushFrame(frame *frame) {
	if len(m.callStack) > 0 {
		m.frame.ip = m.ip
	}
	m.callStack = append(m.callStack, frame)
	m.frame = frame
	if m.frame.fn != nil {
		m.instrs = m.frame.fn.instrs
		m.ip = m.frame.ip
	} else {
		m.instrs = nil
	}
}

func (m *VM) popFrame() {
	f := m.callStack[len(m.callStack)-1]
	*f = frame{}
	m.framePool = append(m.framePool, f)

	m.callStack[len(m.callStack)-1] = nil
	m.callStack = m.callStack[:len(m.callStack)-1]

	if len(m.callStack) > 0 {
		m.frame = m.callStack[len(m.callStack)-1]
		if m.frame.fn != nil {
			m.instrs = m.frame.fn.instrs
			m.ip = m.frame.ip
		}
	}
}

func (m *VM) runFrame(frame *frame) (err error) {
	m.pushFrame(frame)

	defer func() {
		for i := len(m.frame.defers) - 1; i >= 0; i-- {
			deferred := m.frame.defers[i]
			_, derr := m.Call(deferred, nil, 0)
			if derr != nil {
				if err == nil {
					// TODO: combine errors
					err = derr
				}
			}
		}
		m.popFrame()
	}()

	for {
		err := m.resume()
		if err == nil {
			return nil
		}

		// Update the current frame's ip so that the stack trace created by
		// wrapRuntimeError will reflect the current position.
		m.frame.ip = m.ip
		rerr := wrapRuntimeError(m, &err)

		if len(m.frame.tryCatches) == 0 {
			return err
		}

		tryCatch := m.frame.tryCatches[len(m.frame.tryCatches)-1]
		m.frame.tryCatches = m.frame.tryCatches[:len(m.frame.tryCatches)-1]
		m.ip = tryCatch.CatchAddr
		m.sp = m.frame.bp + m.frame.nLocals + 1
		m.stack[m.sp-1] = rerr
	}
}

func (m *VM) resume() (err error) {
	for {
		instr := m.instrs[m.ip]

		switch instr.op {
		case OpNop:

		case OpJump:
			m.ip = int(instr.operand1) - 1

		case OpJumpIfTrue:
			v := m.stack[m.sp-1]
			m.sp--
			b := CoerceToBool(v)
			if b {
				m.ip = int(instr.operand1) - 1
			}

		case OpJumpIfFalse:
			v := m.stack[m.sp-1]
			m.sp--
			b := CoerceToBool(v)
			if !b {
				m.ip = int(instr.operand1) - 1
			}

		case OpDup:
			m.stack[m.sp] = m.stack[m.sp-1]
			m.sp++

		case OpPop:
			m.sp -= int(instr.operand1)

		case OpCall:
			nret := int(instr.operand2)
			narg := int(instr.operand1 & 0x7fffffff)

			expand := (instr.operand1 & 0x80000000) != 0
			if expand {
				if narg == 0 {
					return fmt.Errorf("assert: zero arguments in expansion")
				}
				lastArg := m.stack[m.sp-1]
				if lastArg != nil {
					arr, ok := lastArg.(*Array)
					if !ok {
						return fmt.Errorf("cannot expand %v argument", TypeName(lastArg))
					}
					copy(m.stack[m.sp-1:], arr.array)
					m.sp += len(arr.array) - 1
					narg += len(arr.array) - 1
				} else {
					m.sp--
					narg--
				}
			}

			callable := m.stack[m.sp-narg-1]
			rsp := m.sp - narg - 1 + nret
			err = m.call(callable, narg, nret)
			if err != nil {
				return err
			}
			if nret > 0 {
				rets := m.stack[m.sp-nret : m.sp]
				copy(m.stack[rsp-nret:], rets)
			}
			m.sp = rsp

		case OpNil:
			m.stack[m.sp] = nil
			m.sp++

		case OpNewClosure:
			capN := int(instr.operand2)
			fn := int(instr.operand1)

			var caps []ValueRef
			if capN > 0 {
				caps = make([]ValueRef, capN)
				for i, capture := range m.stack[m.sp-capN : m.sp] {
					caps[i] = capture.(ValueRef)
				}
				m.sp -= capN
			}

			closure := &Closure{
				fn:   &m.program.fns[fn],
				caps: caps,
			}
			m.stack[m.sp] = closure
			m.sp++

		case OpNewIter:
			capN := int(instr.operand2)
			fn := int(instr.operand1 & 0x00ffffff)
			iterNRet := int(instr.operand1 & 0xff000000 >> 24)

			var caps []ValueRef
			if capN > 0 {
				caps = make([]ValueRef, capN)
				for i, capture := range m.stack[m.sp-capN : m.sp] {
					caps[i] = capture.(ValueRef)
				}
				m.sp -= capN
			}

			iter := &Iterator{
				fn:       &m.program.fns[fn],
				captures: caps,
				iterNRet: iterNRet,
			}
			iter.stack = iter.preAllocStack[:]

			m.stack[m.sp] = iter
			m.sp++

		case OpNewInt:
			m.stack[m.sp] = NewInt(int64(instr.operand1))
			m.sp++

		case OpNewBool:
			m.stack[m.sp] = NewBool(instr.operand1 != 0)
			m.sp++

		case OpNewObject:
			m.stack[m.sp] = NewObject()
			m.sp++

		case OpNewArray:
			m.stack[m.sp] = NewArray()
			m.sp++

		case OpLoadGlobal:
			m.stack[m.sp] = m.globals[int(instr.operand1)]
			m.sp++

		case OpLoadGlobalRef:
			m.stack[m.sp] = ValueRef{&m.globals[int(instr.operand1)]}
			m.sp++

		case OpLoadLocal:
			l := m.stack[m.frame.bp+int(instr.operand1)]
			if ref, ok := l.(ValueRef); ok {
				l = *ref.Ref
			}
			m.stack[m.sp] = l
			m.sp++

		case OpLoadLocalRef:
			ref, ok := m.stack[m.frame.bp+int(instr.operand1)].(ValueRef)
			if ok {
				m.stack[m.sp] = ref
			} else {
				m.stack[m.sp] = ValueRef{&m.stack[m.frame.bp+int(instr.operand1)]}
			}
			m.sp++

		case OpCaptureLocal:
			l := m.stack[m.frame.bp+int(instr.operand1)]
			if _, ok := l.(ValueRef); !ok {
				ref := ValueRef{Ref: new(Value)}
				*ref.Ref = l
				m.stack[m.frame.bp+int(instr.operand1)] = ref
				l = ref
			}
			m.stack[m.sp] = l
			m.sp++

		case OpLoadArg:
			idx := int(instr.operand1)
			if idx < m.frame.nArg {
				a := m.stack[m.frame.bp-m.frame.nArg+idx]
				if ref, ok := a.(ValueRef); ok {
					a = *ref.Ref
				}
				m.stack[m.sp] = a
			} else {
				m.stack[m.sp] = nil
			}
			m.sp++

		case OpLoadArgRef:
			idx := int(instr.operand1)
			if idx >= m.frame.nArg {
				// TODO: min arg count
				return fmt.Errorf(
					"cannot assign to arg because it was not provided by caller")
			}
			ref, ok := m.stack[m.frame.bp-m.frame.nArg+idx].(ValueRef)
			if ok {
				m.stack[m.sp] = ref
			} else {
				m.stack[m.sp] = ValueRef{
					&m.stack[m.frame.bp-m.frame.nArg+idx]}
			}
			m.sp++

		case OpCaptureArg:
			a := m.stack[m.frame.bp-m.frame.nArg+int(instr.operand1)]
			if _, ok := a.(ValueRef); !ok {
				ref := ValueRef{Ref: new(Value)}
				*ref.Ref = a
				m.stack[m.frame.bp-m.frame.nArg+int(instr.operand1)] = ref
				a = ref
			}
			m.stack[m.sp] = a
			m.sp++

		case OpLoadCapture:
			m.stack[m.sp] = *m.frame.caps[int(instr.operand1)].Ref
			m.sp++

		case OpLoadCaptureRef:
			m.stack[m.sp] = m.frame.caps[int(instr.operand1)]
			m.sp++

		case OpLoadFn:
			m.stack[m.sp] = &m.program.fns[int(instr.operand1)]
			m.sp++

		case OpLoadNativeFn:
			m.stack[m.sp] = m.program.extFns[int(instr.operand1)]
			m.sp++

		case OpLoadLiteral:
			m.stack[m.sp] = m.program.literals[int(instr.operand1)]
			m.sp++

		case OpEvalBinOp:
			op := BinOp(instr.operand1)
			operand1 := m.stack[m.sp-2]
			operand2 := m.stack[m.sp-1]
			res, err := EvalBinOp(op, operand1, operand2)
			if err != nil {
				return err
			}
			m.stack[m.sp-2] = res
			m.sp--

		case OpNot:
			term := m.stack[m.sp-1]
			m.stack[m.sp-1] = NewBool(!CoerceToBool(term))

		case OpUnaryMinus:
			term := m.stack[m.sp-1]
			if term == nil {
				return errors.New("value is nil")
			}
			res, err := term.EvalUnaryMinus()
			if err != nil {
				return err
			}
			m.stack[m.sp-1] = res

		case OpObjectPutNoPop:
			obj := m.stack[m.sp-3].(*Object)
			key := m.stack[m.sp-2]
			val := m.stack[m.sp-1]
			obj.Put(key, val)
			m.sp -= 2

		case OpObjectGet:
			objRaw := m.stack[m.sp-2]
			key := m.stack[m.sp-1]
			if objRaw == nil {
				m.stack[m.sp-2] = nil
			} else {
				indexable, ok := objRaw.(Indexable)
				if !ok {
					return fmt.Errorf("type %v is not indexable", TypeName(objRaw))
				}
				value, err := indexable.Index(key)
				if err != nil {
					return err
				}
				m.stack[m.sp-2] = value
			}
			m.sp--

		case OpObjectGetRef:
			objRaw := m.stack[m.sp-2]
			key := m.stack[m.sp-1]
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
			m.stack[m.sp-2] = valueRef
			m.sp--

		case OpArrayAppendNoPop:
			array := m.stack[m.sp-2].(*Array)
			value := m.stack[m.sp-1]
			array.Push(value)
			m.sp--

		case OpRet:
			nret := m.sp - (m.frame.bp + m.frame.nLocals)
			if m.frame.nRet > nret {
				return fmt.Errorf(
					"function was expected to return at least %v values, "+
						"but returned only %v",
					m.frame.nRet, nret)
			}
			if nret > m.frame.nRet {
				m.sp -= (nret - m.frame.nRet)
			}
			return nil

		case OpStore:
			count := int(instr.operand1)
			for i := 0; i < count; i++ {
				rval := m.stack[m.sp-(count*2-i)].(ValueRef)
				val := m.stack[m.sp-(count-i)]
				*rval.Ref = val
			}
			m.sp -= count * 2

		case OpInitCallFrame:
			m.frame.nLocals = int(instr.operand1)
			m.frame.bp = m.sp
			m.sp += m.frame.nLocals
			for i := m.frame.bp; i < m.sp; i++ {
				m.stack[i] = nil
			}

		case OpMakeIter:
			v := m.stack[m.sp-1]
			switch v := v.(type) {
			case *Iterator:
				// Ready to go.
			case Iterable:
				m.stack[m.sp-1] = v.Iterate()
			default:
				return fmt.Errorf(
					"cannot iterate over value of type %q", TypeName(v))
			}

		case OpBeginTry:
			m.frame.tryCatches = append(m.frame.tryCatches, tryCatch{
				CatchAddr: int(instr.operand1),
			})

		case OpEndTry:
			m.frame.tryCatches = m.frame.tryCatches[:len(m.frame.tryCatches)-1]
			m.ip = int(instr.operand1) - 1

		case OpSwap:
			i1 := m.sp - 1
			i2 := i1 - int(instr.operand1)
			m.stack[i1], m.stack[i2] = m.stack[i2], m.stack[i1]

		case OpThrow:
			errVal := m.stack[m.sp-1]
			m.sp--
			err, ok := errVal.(*RuntimeError)
			if !ok {
				err = &RuntimeError{ErrValue: errVal}
			}
			return err

		case OpDefer:
			deferClosure := m.stack[m.sp-1].(*Closure)
			m.sp--
			m.frame.defers = append(m.frame.defers, deferClosure)

		case OpNext:
			jumpTo := int(instr.operand1)
			n := int(instr.operand2)
			hasRaw := m.stack[m.sp-1-n]
			has, ok := hasRaw.(Bool)
			if !ok {
				return fmt.Errorf(
					"iterator's first return value must be a bool; but it was %q",
					TypeName(hasRaw))
			}
			if has.Bool() {
				for i := 0; i < n; i++ {
					rval := m.stack[m.sp-1-(n*2-i)].(ValueRef)
					val := m.stack[m.sp-(n-i)]
					*rval.Ref = val
				}
			} else {
				m.ip = jumpTo - 1
			}
			m.sp -= n*2 + 1

		case OpSlice:
			target := m.stack[m.sp-3]
			begin := m.stack[m.sp-2]
			end := m.stack[m.sp-1]

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
			m.stack[m.sp-3] = res
			m.sp -= 2

		case OpIterYield:
			nret := m.sp - (m.frame.bp + m.frame.nLocals)
			if m.frame.nRet > nret {
				return fmt.Errorf(
					"iterator was expected to yield at least %v values, "+
						"but yielded only %v",
					m.frame.nRet, nret)
			}

			iter := m.frame.iter
			iter.tryCatches = m.frame.tryCatches
			iter.defers = m.frame.defers
			iter.nlocals = m.frame.nLocals
			iter.ip = m.ip + 1

			if nret > m.frame.nRet {
				m.sp -= (nret - m.frame.nRet)
			}

			return nil

		case OpIterRet:
			m.frame.iter.ip = -1
			m.stack[m.sp] = False
			for i := 1; i < m.frame.nRet; i++ {
				m.stack[m.sp+i] = nil
			}
			m.sp += m.frame.nRet
			return nil

		default:
			panic("invalid instruction")
		}

		m.ip++
	}
}

func (m *VM) GetStackInfo() []FrameInfo {
	stack := make([]FrameInfo, 0, len(m.callStack))
	for i := len(m.callStack) - 1; i >= 0; i-- {
		stack = append(stack, m.getFrameInfo(m.callStack[i]))
	}
	return stack
}

func (m *VM) GetNArg() int {
	if len(m.callStack) < 2 {
		return 0
	}
	return m.callStack[len(m.callStack)-2].nArg
}

func (m *VM) GetArgs() []Value {
	if len(m.callStack) < 2 {
		return nil
	}
	f := m.callStack[len(m.callStack)-2]
	args := m.stack[f.bp-f.nArg : f.bp]
	return args
}

func (m *VM) getFrameInfo(frame *frame) FrameInfo {
	if frame.fn != nil {
		loc := m.getLocation(frame.fn, frame.ip)
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

	fn := runtime.FuncForPC(reflect.ValueOf(frame.extFn).Pointer())
	if fn == nil {
		return FrameInfo{
			Filename: "???",
			Line:     0,
			Func:     "???",
		}
	}

	fnName := fn.Name()
	lastSlash := strings.LastIndexByte(fnName, '/')
	if lastSlash != -1 {
		fnName = fnName[lastSlash+1:]
	}

	return FrameInfo{
		Func: fnName,
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

func (m *VM) newFrame() *frame {
	if len(m.framePool) == 0 {
		return &frame{}
	}

	f := m.framePool[len(m.framePool)-1]
	m.framePool = m.framePool[:len(m.framePool)-1]
	return f
}