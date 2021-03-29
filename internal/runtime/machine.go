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
	OpSwapStack
	OpThrow
	OpDefer
)

type Instr struct {
	Op       OpCode
	Operand2 byte
	Operand1 uint16
}

func operandsToWord24(operand1 uint16, operand2 byte) uint32 {
	return uint32(operand1) | (uint32(operand2) << 16)
}

func word24ToOperands(word24 uint32) (operand1 uint16, operand2 byte) {
	operand1 = uint16(word24 & 0xFFFF)
	operand2 = byte((word24 >> 16) & 0xFF)
	return operand1, operand2
}

type FrameInfo struct {
	Filename string
	Line     int
}

var machineContextKey = "runtime.Machine"

func MachineFromContext(ctx context.Context) *Machine {
	return ctx.Value(&machineContextKey).(*Machine)
}

type Machine struct {
	callStack []*frame
	program   *Program
	globals   []Value
	reqParamN int
	frame     *frame
}

func NewMachine(prog *Program) *Machine {
	return &Machine{
		program:   prog,
		globals:   make([]Value, prog.globals),
		reqParamN: prog.reqParamN,
	}
}

func (m *Machine) Run(
	ctx context.Context,
	params map[string]Value,
) error {
	ctx = context.WithValue(ctx, &machineContextKey, m)

	for paramName, paramValue := range params {
		param := m.program.params[paramName]
		if param == nil {
			return fmt.Errorf("invalid parameter: %s", paramName)
		}
		if param.required {
			m.reqParamN--
		}
		m.globals[param.global] = paramValue
	}
	if m.reqParamN != 0 {
		return fmt.Errorf("required parameters not set")
	}

	_, err := m.runFunc(ctx, &m.program.fns[0], nil, nil, 0)

	return err
}

func (m *Machine) runFunc(
	ctx context.Context,
	fn *Fn,
	args []Value,
	captures []ValueRef,
	retN int,
) ([]Value, error) {
	frame := newFrame()
	frame.Fn = fn
	frame.Instrs = fn.instrs
	frame.Args = args
	frame.ExpRetN = retN
	frame.Captures = captures
	return m.runFrame(ctx, frame)
}

func (m *Machine) runCallable(
	ctx context.Context,
	callable Value,
	args []Value,
	expRetN int,
) ([]Value, error) {
	var rets []Value
	var err error
	switch callable := callable.(type) {
	case *Closure:
		if callable.extFn != nil {
			rets, err = callable.extFn(ctx, callable.caps, args, expRetN)
			if err != nil {
				return nil, err
			}
			if len(rets) < expRetN {
				return nil, fmt.Errorf("expected at least %v returned values", expRetN)
			}
		} else {
			rets, err = m.runFunc(ctx, callable.fn, args, callable.caps, expRetN)
			if err != nil {
				return nil, err
			}
		}

	case *Fn:
		rets, err = m.runFunc(ctx, callable, args, nil, expRetN)
		if err != nil {
			return nil, err
		}

	case ExternFn:
		rets, err = callable(ctx, nil, args, expRetN)
		if err != nil {
			return nil, err
		}
		if len(rets) < expRetN {
			return nil, fmt.Errorf("expected at least %v returned values", expRetN)
		}

	default:
		if callable == nil {
			return nil, ErrCannotCallNil
		}
		return nil, fmt.Errorf("cannot call type %q", callable.Type())
	}
	return rets, nil
}

func (m *Machine) runFrame(ctx context.Context, frame *frame) (rets []Value, err error) {
	m.callStack = append(m.callStack, frame)
	m.frame = frame

	defer func() {
		for i := len(m.frame.Defers) - 1; i >= 0; i-- {
			deferred := m.frame.Defers[i]
			_, derr := m.runCallable(ctx, deferred, nil, 0)
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
		ret, err := m.resume(ctx)
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

		if len(m.frame.TryCatches) == 0 {
			return nil, err
		}

		tryCatch := m.frame.TryCatches[len(m.frame.TryCatches)-1]
		m.frame.TryCatches = m.frame.TryCatches[:len(m.frame.TryCatches)-1]
		m.frame.IP = tryCatch.CatchAddr
		m.frame.Stack = append(m.frame.Stack[:0], rerr)
	}
}

func (m *Machine) resume(ctx context.Context) (ret []Value, err error) {
	for {
		instr := m.frame.Instrs[m.frame.IP]

		switch instr.Op {
		case OpNop:

		case OpJump:
			m.frame.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1

		case OpJumpIfTrue:
			v := coerceToBool(m.pop())
			if v {
				m.frame.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpJumpIfFalse:
			v := coerceToBool(m.pop())
			if !v {
				m.frame.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpDup:
			v := m.peek(0)
			m.push(v)

		case OpPop:
			m.discardN(int(instr.Operand1))

		case OpCall:
			expRetN := int(instr.Operand2)
			argN := int(instr.Operand1)
			args := make([]Value, argN)
			copy(args, m.popN(argN))
			callable := m.pop()
			rets, err := m.runCallable(ctx, callable, args, expRetN)
			if err != nil {
				return nil, err
			}
			m.frame.Stack = append(m.frame.Stack, rets[:expRetN]...)

		case OpNewClosure:
			capN := int(instr.Operand2)
			fn := int(instr.Operand1)
			caps := make([]ValueRef, capN)
			for i, capture := range m.popN(capN) {
				caps[i] = capture.(ValueRef)
			}
			closure := &Closure{
				fn:   &m.program.fns[fn],
				caps: caps,
			}
			m.push(closure)

		case OpNewInt:
			m.push(NewInt(int64(instr.Operand1)))

		case OpNewBool:
			m.push(NewBool(instr.Operand1 != 0))

		case OpNewObject:
			m.push(NewObject())

		case OpNewArray:
			m.push(NewArray())

		case OpLoadGlobal:
			m.push(m.globals[int(instr.Operand1)])

		case OpLoadGlobalRef:
			m.push(ValueRef{&m.globals[int(instr.Operand1)]})

		case OpLoadLocal:
			m.push(m.frame.Locals[int(instr.Operand1)])

		case OpLoadLocalRef:
			m.push(ValueRef{&m.frame.Locals[int(instr.Operand1)]})

		case OpLoadArg:
			m.push(m.frame.Args[int(instr.Operand1)])

		case OpLoadArgRef:
			m.push(ValueRef{&m.frame.Args[int(instr.Operand1)]})

		case OpLoadCapture:
			m.push(*m.frame.Captures[int(instr.Operand1)].Ref)

		case OpLoadCaptureRef:
			m.push(m.frame.Captures[int(instr.Operand1)])

		case OpLoadFn:
			m.push(&m.program.fns[int(instr.Operand1)])

		case OpLoadExternFn:
			m.push(m.program.extFns[int(instr.Operand1)])

		case OpLoadLiteral:
			m.push(m.program.literals[int(instr.Operand1)])

		case OpEvalBinOp:
			operand2 := m.pop()
			operand1 := m.pop()
			op := BinOp(instr.Operand1)
			res, err := EvalBinOp(op, operand1, operand2)
			if err != nil {
				return nil, err
			}
			m.push(res)

		case OpNot:
			term := m.pop()
			m.push(NewBool(!coerceToBool(term)))

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
					return nil, fmt.Errorf("Value type %v is not indexable", objRaw.Type())
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
					return nil, fmt.Errorf("Value type %v is not indexable", objRaw.Type())
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
			array.Append(value)

		case OpRet:
			if m.frame.ExpRetN > len(m.frame.Stack) {
				return nil, fmt.Errorf("error")
			}
			rets := m.frame.Stack[:m.frame.ExpRetN]
			return rets, nil

		case OpStore:
			count := int(instr.Operand1)
			for i := 0; i < count; i++ {
				rval := m.peek(count*2 - 1 - i).(ValueRef)
				val := m.peek(count - 1 - i)
				*rval.Ref = val
			}
			m.discardN(count * 2)

		case OpInitCallFrame:
			localN := instr.Operand1
			m.frame.Init(int(localN))

		case OpMakeIter:
			v := m.peek(0)
			switch v := v.(type) {
			case *Closure:
				// Ready to go.
			case Enumerable:
				m.pop()
				m.push(v.Enumerate())
			default:
				return nil, fmt.Errorf("Cannot iterate over value of type %q", v.Type())
			}

		case OpBeginTry:
			m.frame.TryCatches = append(m.frame.TryCatches, tryCatch{
				CatchAddr: int(operandsToWord24(instr.Operand1, instr.Operand2)),
			})

		case OpEndTry:
			m.frame.TryCatches = m.frame.TryCatches[:len(m.frame.TryCatches)-1]
			m.frame.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1

		case OpSwapStack:
			i1 := len(m.frame.Stack) - 1
			i2 := i1 - int(instr.Operand1)
			t := m.frame.Stack[i2]
			m.frame.Stack[i2] = m.frame.Stack[i1]
			m.frame.Stack[i1] = t

		case OpThrow:
			errVal := m.pop()
			err, ok := errVal.(*RuntimeError)
			if !ok {
				err = &RuntimeError{ErrValue: errVal}
			}
			return nil, err

		case OpDefer:
			deferClosure := m.pop().(*Closure)
			m.frame.Defers = append(m.frame.Defers, deferClosure)

		default:
			panic("invalid instruction")
		}

		m.frame.IP++
	}
}

func (m *Machine) GetDebugStack() []FrameInfo {
	stack := make([]FrameInfo, 0, len(m.callStack))
	for i := len(m.callStack) - 1; i >= 0; i-- {
		stack = append(stack, m.getDebugFrame(m.callStack[i]))
	}
	return stack
}

func (m *Machine) getDebugFrame(frame *frame) FrameInfo {
	loc := m.getLocation(frame.Fn, frame.IP)
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
	return &locs[len(fn.locations)]
}

func (m *Machine) push(v Value) {
	m.frame.Stack = append(m.frame.Stack, v)
}

func (m *Machine) pop() Value {
	l := len(m.frame.Stack)
	r := m.frame.Stack[l-1]
	m.frame.Stack = m.frame.Stack[:l-1]
	return r
}

func (m *Machine) peek(n int) Value {
	return m.frame.Stack[len(m.frame.Stack)-1-n]
}

func (m *Machine) popN(n int) []Value {
	r := m.frame.Stack[len(m.frame.Stack)-n:]
	m.frame.Stack = m.frame.Stack[:len(m.frame.Stack)-n]
	return r
}

func (m *Machine) discardN(n int) {
	m.frame.Stack = m.frame.Stack[:len(m.frame.Stack)-n]
}
