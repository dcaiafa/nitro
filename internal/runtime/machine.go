package runtime

import (
	"context"
	"errors"
	"fmt"
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
	OpLoadExternFn
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
}

type tryCatch struct {
	CatchAddr int
}

type frame struct {
	nRet       int
	nArg       int
	iter       *Iterator
	fn         *Fn
	extFn      ExternFn
	instrs     []Instr
	args       []Value
	caps       []ValueRef
	locals     []Value
	stack      []Value
	tryCatches []tryCatch
	defers     []*Closure
	ip         int
}

type Machine struct {
	ctx       context.Context
	program   *Program
	globals   []Value
	callStack []*frame
	frame     *frame
	userData  map[interface{}]interface{}
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
	_, err := m.runFunc(&m.program.fns[0], nil, nil, 0)
	return err
}

func (m *Machine) runFunc(
	fn *Fn,
	args []Value,
	captures []ValueRef,
	retN int,
) ([]Value, error) {
	frame := &frame{}
	frame.fn = fn
	frame.instrs = fn.instrs
	frame.args = args
	frame.nRet = retN
	frame.caps = captures
	return m.runFrame(frame)
}

func (m *Machine) callExtFn(extFn ExternFn, args []Value, caps []ValueRef, retN int) ([]Value, error) {
	rets, err := extFn(m, caps, args, retN)
	if err != nil {
		return nil, err
	}
	if len(rets) < retN {
		return nil, fmt.Errorf("expected at least %v returned values", retN)
	}
	return rets, nil
}

func (m *Machine) Call(
	callable Value,
	args []Value,
	expRetN int,
) ([]Value, error) {
	switch callable := callable.(type) {
	case *Closure:
		if callable.extFn != nil {
			return m.callExtFn(
				callable.extFn,
				copyArgs(args, 0),
				callable.caps,
				expRetN)
		} else {
			return m.runFrame(&frame{
				fn:   callable.fn,
				args: copyArgs(args, callable.fn.minArgs),
				caps: callable.caps,
				nArg: len(args),
				nRet: expRetN,
			})
		}

	case *Iterator:
		if callable.extFn != nil {
			return m.callExtFn(
				callable.extFn,
				copyArgs(args, 0),
				callable.captures,
				expRetN)
		} else {
			if callable.ip == -1 {
				rets := make([]Value, expRetN)
				if expRetN > 0 {
					rets[0] = NewBool(false)
				}
				return rets, nil
			}

			return m.runFrame(&frame{
				fn:         callable.fn,
				iter:       callable,
				caps:       callable.captures,
				locals:     callable.locals,
				tryCatches: callable.tryCatches,
				defers:     callable.defers,
				nRet:       expRetN,
				ip:         callable.ip,
			})
		}

	case *Fn:
		return m.runFrame(&frame{
			fn:   callable,
			args: copyArgs(args, callable.minArgs),
			nArg: len(args),
			nRet: expRetN,
		})

	case ExternFn:
		return m.callExtFn(callable, args, nil, expRetN)

	default:
		if callable == nil {
			return nil, ErrCannotCallNil
		}
		return nil, fmt.Errorf("cannot call type %q", TypeName(callable))
	}
}

func (m *Machine) runFrame(frame *frame) (rets []Value, err error) {
	frame.instrs = frame.fn.instrs

	m.callStack = append(m.callStack, frame)
	m.frame = frame

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

		m.callStack[len(m.callStack)-1] = nil
		m.callStack = m.callStack[:len(m.callStack)-1]
		if len(m.callStack) > 0 {
			m.frame = m.callStack[len(m.callStack)-1]
		}
	}()

	for {
		ret, err := m.resume()
		if err == nil {
			return ret, nil
		}

		rerr, ok := err.(*RuntimeError)
		if !ok {
			rerr = &RuntimeError{
				Err: err,
			}
		}

		if rerr.Stack == nil {
			rerr.Stack = m.GetDebugStack()
		}

		err = rerr

		if len(m.frame.tryCatches) == 0 {
			return nil, err
		}

		tryCatch := m.frame.tryCatches[len(m.frame.tryCatches)-1]
		m.frame.tryCatches = m.frame.tryCatches[:len(m.frame.tryCatches)-1]
		m.frame.ip = tryCatch.CatchAddr
		m.frame.stack = append(m.frame.stack[:0], rerr)
	}
}

func (m *Machine) resume() (ret []Value, err error) {
	for {
		instr := m.frame.instrs[m.frame.ip]

		switch instr.op {
		case OpNop:

		case OpJump:
			m.frame.ip = int(instr.operand1) - 1

		case OpJumpIfTrue:
			v := CoerceToBool(m.pop())
			if v {
				m.frame.ip = int(instr.operand1) - 1
			}

		case OpJumpIfFalse:
			v := CoerceToBool(m.pop())
			if !v {
				m.frame.ip = int(instr.operand1) - 1
			}

		case OpDup:
			v := m.peek(0)
			m.push(v)

		case OpPop:
			m.discardN(int(instr.operand1))

		case OpCall:
			nret := int(instr.operand2)
			narg := int(instr.operand1 & 0x7fffffff)
			args := m.popN(narg)

			expand := (instr.operand1 & 0x80000000) != 0
			if expand {
				if len(args) == 0 {
					return nil, errors.New("OpCall: expand bit is set, but no args")
				}
				lastArg := args[len(args)-1]
				args = args[:len(args)-1]
				if lastArg != nil {
					arr, ok := lastArg.(*Array)
					if !ok {
						return nil, fmt.Errorf(
							"cannot expand %q argument",
							TypeName(lastArg))
					}
					for i := 0; i < arr.Len(); i++ {
						args = append(args, arr.Get(i))
					}
				}
			}

			callable := m.pop()
			rets, err := m.Call(callable, args, nret)
			if err != nil {
				return nil, err
			}
			m.frame.stack = append(m.frame.stack, rets[:nret]...)

		case OpNil:
			m.push(nil)

		case OpNewClosure:
			capN := int(instr.operand2)
			fn := int(instr.operand1)
			caps := make([]ValueRef, capN)
			for i, capture := range m.popN(capN) {
				caps[i] = capture.(ValueRef)
			}
			closure := &Closure{
				fn:   &m.program.fns[fn],
				caps: caps,
			}
			m.push(closure)

		case OpNewIter:
			capN := int(instr.operand2)
			fn := int(instr.operand1 & 0x00ffffff)
			iterNRet := int(instr.operand1 & 0xff000000 >> 24)
			caps := make([]ValueRef, capN)
			for i, capture := range m.popN(capN) {
				caps[i] = capture.(ValueRef)
			}
			iter := &Iterator{
				fn:       &m.program.fns[fn],
				captures: caps,
				iterNRet: iterNRet,
			}
			m.push(iter)

		case OpNewInt:
			m.push(NewInt(int64(instr.operand1)))

		case OpNewBool:
			m.push(NewBool(instr.operand1 != 0))

		case OpNewObject:
			m.push(NewObject())

		case OpNewArray:
			m.push(NewArray())

		case OpLoadGlobal:
			m.push(m.globals[int(instr.operand1)])

		case OpLoadGlobalRef:
			m.push(ValueRef{&m.globals[int(instr.operand1)]})

		case OpLoadLocal:
			m.push(m.frame.locals[int(instr.operand1)])

		case OpLoadLocalRef:
			m.push(ValueRef{&m.frame.locals[int(instr.operand1)]})

		case OpLoadArg:
			m.push(m.frame.args[int(instr.operand1)])

		case OpLoadArgRef:
			m.push(ValueRef{&m.frame.args[int(instr.operand1)]})

		case OpLoadCapture:
			m.push(*m.frame.caps[int(instr.operand1)].Ref)

		case OpLoadCaptureRef:
			m.push(m.frame.caps[int(instr.operand1)])

		case OpLoadFn:
			m.push(&m.program.fns[int(instr.operand1)])

		case OpLoadExternFn:
			m.push(m.program.extFns[int(instr.operand1)])

		case OpLoadLiteral:
			m.push(m.program.literals[int(instr.operand1)])

		case OpEvalBinOp:
			operand2 := m.pop()
			operand1 := m.pop()
			op := BinOp(instr.operand1)

			res, err := EvalBinOp(op, operand1, operand2)
			if err != nil {
				return nil, err
			}

			m.push(res)

		case OpNot:
			term := m.pop()
			m.push(NewBool(!CoerceToBool(term)))

		case OpUnaryMinus:
			term := m.pop()
			if term == nil {
				return nil, errors.New("value is nil")
			}
			switch term := term.(type) {
			case Int:
				m.push(NewInt(-term.Int64()))
			case Float:
				m.push(NewFloat(-term.Float64()))
			default:
			}

		case OpObjectPutNoPop:
			val := m.pop()
			key := m.pop()
			obj := m.peek(0).(*Object)
			obj.Put(key, val)

		case OpObjectGet:
			member := m.pop()
			objRaw := m.pop()
			if objRaw == nil {
				m.push(nil)
			} else {
				indexable, ok := objRaw.(Indexable)
				if !ok {
					return nil, fmt.Errorf("Value type %v is not indexable", TypeName(objRaw))
				}
				item, err := indexable.Index(member)
				if err != nil {
					return nil, err
				}
				m.push(item)
			}

		case OpObjectGetRef:
			member := m.pop()
			objRaw := m.pop()
			if objRaw == nil {
				m.push(nil)
			} else {
				indexable, ok := objRaw.(Indexable)
				if !ok {
					return nil, fmt.Errorf("Value type %v is not indexable", TypeName(objRaw))
				}
				itemRef, err := indexable.IndexRef(member)
				if err != nil {
					return nil, err
				}
				m.push(itemRef)
			}

		case OpArrayAppendNoPop:
			value := m.pop()
			array := m.peek(0).(*Array)
			array.Push(value)

		case OpRet:
			if m.frame.nRet > len(m.frame.stack) {
				return nil, fmt.Errorf("error")
			}
			rets := m.frame.stack[:m.frame.nRet]
			return rets, nil

		case OpStore:
			count := int(instr.operand1)
			for i := 0; i < count; i++ {
				rval := m.peek(count*2 - 1 - i).(ValueRef)
				val := m.peek(count - 1 - i)
				*rval.Ref = val
			}
			m.discardN(count * 2)

		case OpInitCallFrame:
			localN := instr.operand1
			m.frame.locals = make([]Value, localN)

		case OpMakeIter:
			v := m.peek(0)
			switch v := v.(type) {
			case *Iterator:
				// Ready to go.
			case Iterable:
				m.pop()
				m.push(v.Iterate())
			default:
				return nil, fmt.Errorf(
					"cannot iterate over value of type %q", TypeName(v))
			}

		case OpBeginTry:
			m.frame.tryCatches = append(m.frame.tryCatches, tryCatch{
				CatchAddr: int(instr.operand1),
			})

		case OpEndTry:
			m.frame.tryCatches = m.frame.tryCatches[:len(m.frame.tryCatches)-1]
			m.frame.ip = int(instr.operand1) - 1

		case OpSwap:
			i1 := len(m.frame.stack) - 1
			i2 := i1 - int(instr.operand1)
			t := m.frame.stack[i2]
			m.frame.stack[i2] = m.frame.stack[i1]
			m.frame.stack[i1] = t

		case OpThrow:
			errVal := m.pop()
			err, ok := errVal.(*RuntimeError)
			if !ok {
				err = &RuntimeError{ErrValue: errVal}
			}
			return nil, err

		case OpDefer:
			deferClosure := m.pop().(*Closure)
			m.frame.defers = append(m.frame.defers, deferClosure)

		case OpNext:
			jumpTo := int(instr.operand1)
			argN := int(instr.operand2)
			has, ok := m.peek(argN - 1).(Bool)
			if !ok {
				return nil, fmt.Errorf(
					"iterator's first return value must be Bool; instead it is %q",
					TypeName(m.peek(argN-1)))
			}
			if has.Bool() {
				count := argN - 1
				for i := 0; i < count; i++ {
					rval := m.peek(count*2 - i).(ValueRef)
					val := m.peek(count - 1 - i)
					*rval.Ref = val
				}
			} else {
				m.frame.ip = jumpTo - 1
			}
			m.discardN((argN-1)*2 + 1)

		case OpSlice:
			end := m.pop()
			begin := m.pop()
			target := m.pop()

			sliceable, ok := target.(interface {
				Slice(begin Value, end Value) (Value, error)
			})
			if !ok {
				return nil, fmt.Errorf("cannot slice %q", TypeName(target))
			}

			res, err := sliceable.Slice(begin, end)
			if err != nil {
				return nil, err
			}
			m.push(res)

		case OpIterYield:
			if m.frame.nRet > len(m.frame.stack) {
				return nil, fmt.Errorf("error")
			}

			iter := m.frame.iter
			iter.locals = m.frame.locals
			iter.tryCatches = m.frame.tryCatches
			iter.defers = m.frame.defers
			iter.ip = m.frame.ip + 1

			rets := m.frame.stack[:m.frame.nRet]
			return rets, nil

		case OpIterRet:
			m.frame.iter.ip = -1
			rets := make([]Value, m.frame.nRet)
			if m.frame.nRet > 0 {
				rets[0] = NewBool(false)
			}
			return rets, nil

		default:
			panic("invalid instruction")
		}

		m.frame.ip++
	}
}

func (m *Machine) GetDebugStack() []FrameInfo {
	stack := make([]FrameInfo, 0, len(m.callStack))
	for i := len(m.callStack) - 1; i >= 0; i-- {
		stack = append(stack, m.getDebugFrame(m.callStack[i]))
	}
	return stack
}

func (m *Machine) GetNArg() int {
	return m.callStack[len(m.callStack)-1].nArg
}

func (m *Machine) GetArgs() []Value {
	return m.callStack[len(m.callStack)-1].args
}

func (m *Machine) getDebugFrame(frame *frame) FrameInfo {
	loc := m.getLocation(frame.fn, frame.ip)
	if loc == nil {
		return FrameInfo{
			Filename: "???",
			Line:     0,
		}
	}
	return FrameInfo{
		Filename: m.program.literals[loc.filename].(String).String(),
		Line:     loc.lineNum,
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

func (m *Machine) push(v Value) {
	m.frame.stack = append(m.frame.stack, v)
}

func (m *Machine) pop() Value {
	l := len(m.frame.stack)
	r := m.frame.stack[l-1]
	m.frame.stack = m.frame.stack[:l-1]
	return r
}

func (m *Machine) peek(n int) Value {
	return m.frame.stack[len(m.frame.stack)-1-n]
}

func (m *Machine) popN(n int) []Value {
	r := m.frame.stack[len(m.frame.stack)-n:]
	m.frame.stack = m.frame.stack[:len(m.frame.stack)-n]
	return r
}

func (m *Machine) discardN(n int) {
	m.frame.stack = m.frame.stack[:len(m.frame.stack)-n]
}

func copyArgs(args []Value, minArgs int) []Value {
	n := len(args)
	if n < minArgs {
		n = minArgs
	}
	cargs := make([]Value, n)
	copy(cargs, args)
	return cargs
}
