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
	OpObjectPutNoPop
	OpObjectGet
	OpObjectGetRef
	OpRet
	OpStore
	OpInitCallFrame
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

type Machine struct {
	callFrameFactory CallFrameFactory
	callStack        []*CallFrame
	program          *Program
	globals          []Value
}

func NewMachine() *Machine {
	return &Machine{}
}

func (m *Machine) Run(ctx context.Context, p *Program) error {
	m.program = p
	m.globals = make([]Value, p.globals)
	return m.run(ctx)
}

func (m *Machine) run(ctx context.Context) error {
	ip := 0
	f := m.callFrameFactory.NewCallFrame()
	f.Instrs = m.program.fns[0].instrs

	push := func(v Value) {
		f.Stack = append(f.Stack, v)
	}

	pop := func() Value {
		l := len(f.Stack)
		r := f.Stack[l-1]
		f.Stack = f.Stack[:l-1]
		return r
	}

	peek := func() Value {
		return f.Stack[len(f.Stack)-1]
	}

	popN := func(n int) []Value {
		r := f.Stack[len(f.Stack)-n:]
		f.Stack = f.Stack[:len(f.Stack)-n]
		return r
	}

	for {
		instr := f.Instrs[ip]
		switch instr.Op {
		case OpNop:

		case OpJump:
			ip = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1

		case OpJumpIfTrue:
			v := coerceToBool(pop())
			if v {
				ip = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpJumpIfFalse:
			v := coerceToBool(pop())
			if !v {
				ip = int(operandsToWord24(instr.Operand1, instr.Operand2)) - 1
			}

		case OpDup:
			v := peek()
			push(v)

		case OpPop:
			pop()

		case OpCall:
			expRetN := int(instr.Operand2)
			argN := int(instr.Operand1)
			args := popN(argN)

			switch callable := pop().(type) {
			case *Closure:
				f.IP = ip
				m.callStack = append(m.callStack, f)

				f = m.callFrameFactory.NewCallFrame()
				f.Instrs = callable.Fn.instrs
				f.ExpRetN = expRetN
				f.Args = args
				f.Captures = callable.Captures

				// ip will be 0 after incrementing.
				ip = -1

			case *Fn:
				f.IP = ip
				m.callStack = append(m.callStack, f)

				f = m.callFrameFactory.NewCallFrame()
				f.Instrs = callable.instrs
				f.ExpRetN = expRetN
				f.Args = args

				// ip will be 0 after incrementing.
				ip = -1

			case ExternFn:
				rets, err := callable(ctx, args)
				if err != nil {
					return err
				}
				if len(rets) < expRetN {
					return fmt.Errorf("expected at least %v returned values", expRetN)
				}
				f.Stack = append(f.Stack, rets[:expRetN]...)

			default:
				return fmt.Errorf("cannot call type %q", reflect.TypeOf(callable))
			}

		case OpNewClosure:
			capN := int(instr.Operand2)
			fn := int(instr.Operand1)
			caps := make([]ValueRef, capN)
			for i, capture := range popN(capN) {
				caps[i] = capture.(ValueRef)
			}
			closure := &Closure{
				Fn:       &m.program.fns[fn],
				Captures: caps,
			}
			push(closure)

		case OpNewInt:
			push(Int(instr.Operand1))

		case OpNewBool:
			push(Bool(instr.Operand1 != 0))

		case OpNewObject:
			push(NewObject())

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

		case OpObjectPutNoPop:
			val := pop()
			key := pop()
			obj := peek().(*Object)
			obj.Put(key, val)

		case OpObjectGet:
			member := pop()
			objRaw := pop()
			if objRaw == nil {
				push(nil)
			} else {
				obj, ok := objRaw.(*Object)
				if !ok {
					return errors.New("not an object")
				}
				val := obj.Get(member)
				push(val)
			}

		case OpObjectGetRef:
			member := pop()
			objRaw := pop()
			if objRaw == nil {
				return errors.New("cannot access object member: value is nil")
			}
			obj, ok := objRaw.(*Object)
			if !ok {
				return errors.New("not an object")
			}
			valRef := obj.GetRef(member)
			push(ValueRef{valRef})

		case OpRet:
			if f.ExpRetN > len(f.Stack) {
				return fmt.Errorf("error")
			}
			rets := f.Stack[:f.ExpRetN]
			if len(m.callStack) == 0 {
				return nil
			}

			f = m.callStack[len(m.callStack)-1]
			m.callStack = m.callStack[:len(m.callStack)-1]
			f.Stack = append(f.Stack, rets...)
			ip = f.IP

		case OpStore:
			val := pop()
			rval := pop().(ValueRef)
			*rval.Ref = val

		case OpInitCallFrame:
			localN := instr.Operand1
			f.Init(int(localN))

		default:
			panic("invalid instruction")
		}

		ip++
	}
}

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	switch operand1 := operand1.(type) {
	case Int:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpInt(op, operand1, operand2)
		case Float:
			return evalBinOpFloat(op, Float(operand1), operand2)
		}
	case Float:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpFloat(op, operand1, Float(operand2))
		case Float:
			return evalBinOpFloat(op, operand1, operand2)
		}
		//case string:
	}

	return nil, fmt.Errorf(
		"could not evaluate binary expression: type %v is incompatible with type %v",
		reflect.TypeOf(operand1), reflect.TypeOf(operand2))
}

func evalBinOpInt(op BinOp, operand1, operand2 Int) (Value, error) {
	switch op {
	case BinAdd:
		return operand1 + operand2, nil
	case BinSub:
		return operand1 - operand2, nil
	case BinMult:
		return operand1 * operand2, nil
	case BinDiv:
		return operand1 / operand2, nil
	case BinMod:
		return operand1 % operand2, nil
	case BinLT:
		return Bool(operand1 < operand2), nil
	case BinLE:
		return Bool(operand1 <= operand2), nil
	case BinGT:
		return Bool(operand1 > operand2), nil
	case BinGE:
		return Bool(operand1 >= operand2), nil
	case BinEq:
		return Bool(operand1 == operand2), nil
	case BinNE:
		return Bool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func evalBinOpFloat(op BinOp, operand1, operand2 Float) (Value, error) {
	switch op {
	case BinAdd:
		return operand1 + operand2, nil
	case BinSub:
		return operand1 - operand2, nil
	case BinMult:
		return operand1 * operand2, nil
	case BinDiv:
		return operand1 / operand2, nil
	case BinMod:
		return nil, errors.New("modulo operation not permitted with Float")
	case BinLT:
		return Bool(operand1 < operand2), nil
	case BinLE:
		return Bool(operand1 <= operand2), nil
	case BinGT:
		return Bool(operand1 > operand2), nil
	case BinGE:
		return Bool(operand1 >= operand2), nil
	case BinEq:
		return Bool(operand1 == operand2), nil
	case BinNE:
		return Bool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func coerceToBool(v Value) bool {
	if v == nil {
		return false
	}
	switch v := v.(type) {
	case Bool:
		return bool(v)

	case Int:
		return v != 0

	case Float:
		return v != 0

	case String:
		return len(v) != 0

	default:
		return true
	}
}
