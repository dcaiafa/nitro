package runtime

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

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

type OpCode byte

const (
	OpNop OpCode = iota
	OpJump
	OpJumpIfTrue
	OpJumpIfFalse
	OpDup
	OpPop
	OpCall
	OpNil
	OpNewClosure
	OpNewInt
	OpNewBool
	OpNewObject
	OpNewArray
	OpLoadGlobal
	OpLoadGlobalRef
	OpLoadLocal
	OpLoadLocalRef
	OpLoadArg
	OpLoadArgRef
	OpLoadCapture
	OpLoadCaptureRef
	OpLoadFn
	OpLoadNativeFn
	OpLoadLiteral
	OpEvalBinOp
	OpNot
	OpUnaryMinus
	OpObjectPutNoPop
	OpObjectGet
	OpObjectGetRef
	OpArrayAppendNoPop
	OpRet
	OpStore
	OpInitCallFrame
	OpMakeIter
	OpBeginTry
	OpEndTry
	OpSwap
	OpThrow
	OpDefer
	OpNext
	OpSlice
	OpIterYield
	OpIterRet
	OpNewIter
)

type Instr struct {
	op       OpCode
	operand1 uint32
	operand2 uint16
}

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

type Machine struct {
	ctx       context.Context
	userData  map[interface{}]interface{}
	program   *Program
	globals   []Value
	callStack []*frame
	frame     *frame
	stack     [1000]Value
	instrs    []Instr
	sp        int
	ip        int
}

func NewMachine(ctx context.Context, prog *Program) *Machine {
	return &Machine{
		ctx:      ctx,
		program:  prog,
		globals:  make([]Value, prog.globals),
		userData: make(map[interface{}]interface{}),
	}
}

func (m *Machine) Context() context.Context {
	return m.ctx
}

func (m *Machine) SetUserData(key, value interface{}) {
	m.userData[key] = value
}

func (m *Machine) GetUserData(key interface{}) interface{} {
	return m.userData[key]
}

func (m *Machine) SetParam(n string, v Value) error {
	param := m.program.params[n]
	if param == nil {
		return fmt.Errorf("invalid parameter: %s", n)
	}
	m.globals[param.global] = v
	return nil
}

func (m *Machine) Run() error {
	return m.runFrame(&frame{
		fn: &m.program.fns[0],
	})
}

func (m *Machine) Call(callable Value, args []Value, nret int) ([]Value, error) {
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

func (m *Machine) callExtFn(
	extFn NativeFn,
	caps []ValueRef,
	narg int,
	nret int,
) (err error) {
	frame := &frame{
		nRet:  nret,
		nArg:  narg,
		extFn: extFn,
		caps:  caps,
		bp:    m.sp,
	}

	m.pushFrame(frame)
	defer m.popFrame()

	defer func() {
		if err != nil {
			rerr, ok := err.(*RuntimeError)
			if !ok {
				rerr = &RuntimeError{
					Err: err,
				}
			}
			if rerr.Stack == nil {
				rerr.Stack = m.GetStackInfo()
			}
			err = rerr
		}
	}()

	args := m.stack[m.sp-narg : m.sp]

	rets, err := extFn(m, caps, args, nret)
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

func (m *Machine) call(callable Value, narg int, nret int) error {
	switch callable := callable.(type) {
	case *Closure:
		if callable.extFn != nil {
			return m.callExtFn(callable.extFn, callable.caps, narg, nret)
		} else {
			return m.runFrame(&frame{
				fn:   callable.fn,
				caps: callable.caps,
				nArg: narg,
				nRet: nret,
			})
		}

	case *Iterator:
		if callable.extFn != nil {
			return m.callExtFn(callable.extFn, callable.captures, narg, nret)
		} else {
			if callable.ip == -1 {
				m.stack[m.sp] = False
				for i := 0; i < m.frame.nRet; i++ {
					m.stack[m.sp+1+i] = nil
				}
				m.sp += m.frame.nRet + 1
				return nil
			}

			return m.runFrame(&frame{
				fn:         callable.fn,
				iter:       callable,
				caps:       callable.captures,
				tryCatches: callable.tryCatches,
				defers:     callable.defers,
				nRet:       nret,
				ip:         callable.ip,
			})
		}

	case *Fn:
		return m.runFrame(&frame{
			fn:   callable,
			nArg: narg,
			nRet: nret,
		})

	case NativeFn:
		return m.callExtFn(callable, nil, narg, nret)

	default:
		if callable == nil {
			return ErrCannotCallNil
		}
		return fmt.Errorf("cannot call type %q", TypeName(callable))
	}
}

func (m *Machine) pushFrame(frame *frame) {
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

func (m *Machine) popFrame() {
	m.callStack[len(m.callStack)-1] = nil
	m.callStack = m.callStack[:len(m.callStack)-1]
	if len(m.callStack) > 0 {
		m.frame = m.callStack[len(m.callStack)-1]
		m.instrs = m.frame.fn.instrs
		m.ip = m.frame.ip
	}
}

func (m *Machine) runFrame(frame *frame) (err error) {
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

		rerr, ok := err.(*RuntimeError)
		if !ok {
			rerr = &RuntimeError{
				Err: err,
			}
		}

		if rerr.Stack == nil {
			rerr.Stack = m.GetStackInfo()
		}

		err = rerr

		if len(m.frame.tryCatches) == 0 {
			return err
		}

		tryCatch := m.frame.tryCatches[len(m.frame.tryCatches)-1]
		m.frame.tryCatches = m.frame.tryCatches[:len(m.frame.tryCatches)-1]
		m.ip = tryCatch.CatchAddr
		m.sp = m.frame.bp + m.frame.nLocals
	}
}

func (m *Machine) resume() (err error) {
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
				panic("not implemented")
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
				caps := make([]ValueRef, capN)
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
				caps := make([]ValueRef, capN)
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
			m.stack[m.sp] = m.stack[m.frame.bp+int(instr.operand1)]
			m.sp++

		case OpLoadLocalRef:
			m.stack[m.sp] = ValueRef{&m.stack[m.frame.bp+int(instr.operand1)]}
			m.sp++

		case OpLoadArg:
			m.stack[m.sp] = m.stack[m.frame.bp-m.frame.nArg+int(instr.operand1)]
			m.sp++

		case OpLoadArgRef:
			m.stack[m.sp] = ValueRef{
				&m.stack[m.frame.bp-m.frame.nArg+int(instr.operand1)]}
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
			var res Value
			switch term := term.(type) {
			case Int:
				res = NewInt(-term.Int64())
			case Float:
				res = NewFloat(-term.Float64())
			case Evaluator:
				res, err = term.EvalUnaryMinus()
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf(
					"type %v does not support unary minus operation",
					TypeName(term))
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
			iter.ip = m.frame.ip + 1

			if nret > m.frame.nRet {
				m.sp -= (nret - m.frame.nRet)
			}

			return nil

		case OpIterRet:
			m.frame.iter.ip = -1
			m.stack[m.sp] = False
			for i := 0; i < m.frame.nRet; i++ {
				m.stack[m.sp+1+i] = nil
			}
			m.sp += m.frame.nRet + 1
			return nil

		default:
			panic("invalid instruction")
		}

		m.ip++
	}
}

func (m *Machine) GetStackInfo() []FrameInfo {
	stack := make([]FrameInfo, 0, len(m.callStack))
	for i := len(m.callStack) - 1; i >= 0; i-- {
		stack = append(stack, m.getFrameInfo(m.callStack[i]))
	}
	return stack
}

func (m *Machine) GetNArg() int {
	if len(m.callStack) < 2 {
		return 0
	}
	return m.callStack[len(m.callStack)-2].nArg
}

func (m *Machine) GetArgs() []Value {
	if len(m.callStack) < 2 {
		return nil
	}
	f := m.callStack[len(m.callStack)-2]
	args := m.stack[f.bp-f.nArg : f.bp]
	return args
}

func (m *Machine) getFrameInfo(frame *frame) FrameInfo {
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
		Filename: "<builtin>",
		Line:     0,
		Func:     fnName,
	}
}

func (m *Machine) getLocation(fn *Fn, ip int) *Location {
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
