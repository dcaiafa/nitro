package runtime

import (
	"context"
	"errors"
	"fmt"
	"reflect"
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

type Frame struct {
	Filename string
	Line     int
}

type Machine struct {
	callFrameFactory CallFrameFactory
	callStack        []*CallFrame
	program          *Program
	globals          []Value
	reqParamN        int
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
) (err error) {
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

	f := m.callFrameFactory.NewCallFrame()
	f.Fn = &m.program.fns[0]
	f.Instrs = m.program.fns[0].instrs

	defer func() {
		if err != nil {
			rerr, ok := err.(*RuntimeError)
			if !ok {
				rerr = &RuntimeError{
					Err: err,
				}
			}
			if rerr.Stack == nil {
				rerr.Stack = m.GetDebugStack(f)
			}
			err = rerr
		}
	}()

	push := func(v Value) {
		f.Stack = append(f.Stack, v)
	}

	pop := func() Value {
		l := len(f.Stack)
		r := f.Stack[l-1]
		f.Stack = f.Stack[:l-1]
		return r
	}

	peek := func(n int) Value {
		return f.Stack[len(f.Stack)-1-n]
	}

	popN := func(n int) []Value {
		r := f.Stack[len(f.Stack)-n:]
		f.Stack = f.Stack[:len(f.Stack)-n]
		return r
	}

	discardN := func(n int) {
		f.Stack = f.Stack[:len(f.Stack)-n]
	}

	for {
		instr := f.Instrs[f.IP]
		switch instr.Op {
		case OpNop:

		case OpJump:
			f.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1

		case OpJumpIfTrue:
			v := coerceToBool(pop())
			if v {
				f.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpJumpIfFalse:
			v := coerceToBool(pop())
			if !v {
				f.IP = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpDup:
			v := peek(0)
			push(v)

		case OpPop:
			discardN(int(instr.Operand1))

		case OpCall:
			expRetN := int(instr.Operand2)
			argN := int(instr.Operand1)
			args := make([]Value, argN)
			copy(args, popN(argN))

			switch callable := pop().(type) {
			case *Closure:
				if callable.extFn != nil {
					rets, err := callable.extFn(ctx, callable.caps, args)
					if err != nil {
						return err
					}
					if len(rets) < expRetN {
						return fmt.Errorf("expected at least %v returned values", expRetN)
					}
					f.Stack = append(f.Stack, rets[:expRetN]...)
				} else {
					m.callStack = append(m.callStack, f)

					f = m.callFrameFactory.NewCallFrame()
					f.Fn = callable.fn
					f.Instrs = callable.fn.instrs
					f.ExpRetN = expRetN
					f.Args = args
					f.Captures = callable.caps
					f.IP = -1 // ip will be 0 after incrementing.
				}

			case *Fn:
				m.callStack = append(m.callStack, f)

				f = m.callFrameFactory.NewCallFrame()
				f.Fn = callable
				f.Instrs = callable.instrs
				f.ExpRetN = expRetN
				f.Args = args
				f.IP = -1 // ip will be 0 after incrementing.

			case ExternFn:
				rets, err := callable(ctx, nil, args)
				if err != nil {
					return err
				}
				if len(rets) < expRetN {
					return fmt.Errorf("expected at least %v returned values", expRetN)
				}
				f.Stack = append(f.Stack, rets[:expRetN]...)

			default:
				if callable == nil {
					return ErrCannotCallNil
				}

				return fmt.Errorf("cannot call type %q", callable.Type())
			}

		case OpNewClosure:
			capN := int(instr.Operand2)
			fn := int(instr.Operand1)
			caps := make([]ValueRef, capN)
			for i, capture := range popN(capN) {
				caps[i] = capture.(ValueRef)
			}
			closure := &Closure{
				fn:   &m.program.fns[fn],
				caps: caps,
			}
			push(closure)

		case OpNewInt:
			push(NewInt(int64(instr.Operand1)))

		case OpNewBool:
			push(NewBool(instr.Operand1 != 0))

		case OpNewObject:
			push(NewObject())

		case OpNewArray:
			push(NewArray())

		case OpLoadGlobal:
			push(m.globals[int(instr.Operand1)])

		case OpLoadGlobalRef:
			push(ValueRef{&m.globals[int(instr.Operand1)]})

		case OpLoadLocal:
			push(f.Locals[int(instr.Operand1)])

		case OpLoadLocalRef:
			push(ValueRef{&f.Locals[int(instr.Operand1)]})

		case OpLoadArg:
			push(f.Args[int(instr.Operand1)])

		case OpLoadArgRef:
			push(ValueRef{&f.Args[int(instr.Operand1)]})

		case OpLoadCapture:
			push(*f.Captures[int(instr.Operand1)].Ref)

		case OpLoadCaptureRef:
			push(f.Captures[int(instr.Operand1)])

		case OpLoadFn:
			push(&m.program.fns[int(instr.Operand1)])

		case OpLoadExternFn:
			push(m.program.extFns[int(instr.Operand1)])

		case OpLoadLiteral:
			push(m.program.literals[int(instr.Operand1)])

		case OpEvalBinOp:
			operand2 := pop()
			operand1 := pop()
			op := BinOp(instr.Operand1)
			res, err := EvalBinOp(op, operand1, operand2)
			if err != nil {
				return err
			}
			push(res)

		case OpNot:
			term := pop()
			push(NewBool(!coerceToBool(term)))

		case OpUnaryMinus:
			term := pop()
			if term == nil {
				return errors.New("value is nil")
			}
			switch term := term.(type) {
			case Int:
				push(NewInt(-term.Int64()))
			case Float:
				push(NewFloat(-term.Float64()))
			default:
			}

		case OpObjectPutNoPop:
			val := pop()
			key := pop()
			obj := peek(0).(*Object)
			obj.Put(key, val)

		case OpObjectGet:
			member := pop()
			objRaw := pop()
			if objRaw == nil {
				push(nil)
			} else {
				var v Value
				switch obj := objRaw.(type) {
				case *Object:
					v = obj.Get(member)
				case *Array:
					index, ok := member.(Int)
					if !ok {
						return fmt.Errorf(
							"Cannot index array: index must be Int, but it is %v",
							index.Type())
					}
					v = obj.Get(int(index.Int64()))
				default:
					return fmt.Errorf(
						"Cannot index: allowed types are Object and Array, but got %v",
						objRaw.Type())
				}
				push(v)
			}

		case OpObjectGetRef:
			member := pop()
			objRaw := pop()
			if objRaw == nil {
				return fmt.Errorf("Cannot deref nil value")
			}
			var v *Value
			switch obj := objRaw.(type) {
			case *Object:
				v = obj.GetRef(member)
			case *Array:
				index, ok := member.(Int)
				if !ok {
					return fmt.Errorf(
						"Cannot index array: index must be Int, but it is %v",
						index.Type())
				}
				v = obj.GetRef(int(index.Int64()))
			default:
				return fmt.Errorf(
					"Cannot index: allowed types are Object and Array, but got %v",
					objRaw.Type())
			}
			push(ValueRef{v})

		case OpArrayAppendNoPop:
			value := pop()
			array := peek(0).(*Array)
			array.Append(value)

		case OpRet:
			if len(m.callStack) == 0 {
				return nil
			}
			if f.ExpRetN > len(f.Stack) {
				return fmt.Errorf("error")
			}
			rets := f.Stack[:f.ExpRetN]
			f = m.callStack[len(m.callStack)-1]
			m.callStack = m.callStack[:len(m.callStack)-1]
			f.Stack = append(f.Stack, rets...)

		case OpStore:
			count := int(instr.Operand1)
			for i := 0; i < count; i++ {
				rval := peek(count*2 - 1 - i).(ValueRef)
				val := peek(count - 1 - i)
				*rval.Ref = val
			}
			discardN(count * 2)

		case OpInitCallFrame:
			localN := instr.Operand1
			f.Init(int(localN))

		case OpMakeIter:
			v := peek(0)
			switch v.Kind() {
			case FuncKind:
				// Ready to go.
			case ArrayKind:
				var arr Value = pop().(*Array)
				var next Value = NewInt(0)
				iter := NewClosure(
					arrayIter,
					[]ValueRef{NewValueRef(&arr), NewValueRef(&next)})
				push(iter)
			default:
				return fmt.Errorf("Cannot iterate over value of type %q", v.Type())
			}

		default:
			panic("invalid instruction")
		}

		f.IP++
	}
}

func (m *Machine) GetDebugStack(currentFrame *CallFrame) []Frame {
	stack := make([]Frame, 0, len(m.callStack)+1)
	stack = append(stack, m.getDebugFrame(currentFrame))
	for i := len(m.callStack) - 1; i >= 0; i-- {
		stack = append(stack, m.getDebugFrame(m.callStack[i]))
	}
	return stack
}

func (m *Machine) getDebugFrame(callFrame *CallFrame) Frame {
	loc := m.getLocation(callFrame.Fn, callFrame.IP)
	if loc == nil {
		return Frame{
			Filename: "???",
			Line:     0,
		}
	}
	return Frame{
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

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	switch operand1 := operand1.(type) {
	case Int:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpInt(op, operand1.Int64(), operand2.Int64())
		case Float:
			return evalBinOpFloat(op, float64(operand1.Int64()), operand2.Float64())
		}
	case Float:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpFloat(op, operand1.Float64(), float64(operand2.Int64()))
		case Float:
			return evalBinOpFloat(op, operand1.Float64(), operand2.Float64())
		}

	case String:
		if operand2, ok := operand2.(String); ok {
			return evalBinOpString(op, operand1.String(), operand2.String())
		}
	}

	return nil, fmt.Errorf(
		"could not evaluate binary expression: type %v is incompatible with type %v",
		reflect.TypeOf(operand1), reflect.TypeOf(operand2))
}

func evalBinOpInt(op BinOp, operand1, operand2 int64) (Value, error) {
	switch op {
	case BinAdd:
		return NewInt(operand1 + operand2), nil
	case BinSub:
		return NewInt(operand1 - operand2), nil
	case BinMult:
		return NewInt(operand1 * operand2), nil
	case BinDiv:
		return NewInt(operand1 / operand2), nil
	case BinMod:
		return NewInt(operand1 % operand2), nil
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func evalBinOpFloat(op BinOp, operand1, operand2 float64) (Value, error) {
	switch op {
	case BinAdd:
		return NewFloat(operand1 + operand2), nil
	case BinSub:
		return NewFloat(operand1 - operand2), nil
	case BinMult:
		return NewFloat(operand1 * operand2), nil
	case BinDiv:
		return NewFloat(operand1 / operand2), nil
	case BinMod:
		return nil, errors.New("modulo operation not permitted with Float")
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func evalBinOpString(op BinOp, operand1, operand2 string) (Value, error) {
	switch op {
	case BinAdd:
		return NewString(operand1 + operand2), nil
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		return nil, fmt.Errorf("cannot use this operator with string operands")
	}
}

func coerceToBool(v Value) bool {
	switch v := v.(type) {
	case Bool:
		return v.Bool()
	case Int:
		return v.Int64() != 0
	case Float:
		return v.Float64() != 0
	case String:
		return len(v.String()) != 0
	case *Object:
		return v.Len() != 0
	case *Array:
		return v.Len() != 0
	default:
		return v != nil
	}
}
