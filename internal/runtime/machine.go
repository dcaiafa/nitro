package runtime

import (
	"context"
	"fmt"
	"reflect"
)

type BinOp byte

const (
	BinAdd BinOp = iota
	BinSub
	BinMult
	BinDiv
	BinLT
	BinLE
	BinGT
	BinGE
)

type OpCode byte

const (
	OpNop OpCode = iota
	OpCall
	OpMakeClosure
	OpPushInt
	OpPushGlobal
	OpPushGlobalRef
	OpPushLocal
	OpPushLocalRef
	OpPushArg
	OpPushArgRef
	OpPushFn
	OpPushExternFn
	OpPushLiteral
	OpBinOp
	OpRet
	OpStore
	OpInitCallFrame
)

type Instr struct {
	Op       OpCode
	Operand2 byte
	Operand1 uint16
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
	var (
		ip = 0
		f  = m.callFrameFactory.NewCallFrame()
	)

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

	popN := func(n int) []Value {
		r := f.Stack[len(f.Stack)-n:]
		f.Stack = f.Stack[:len(f.Stack)-n]
		return r
	}

	vcopy := func(v []Value) []Value {
		r := make([]Value, len(v))
		copy(r, v)
		return r
	}

	for {
		instr := f.Instrs[ip]
		switch instr.Op {
		case OpNop:

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

		case OpMakeClosure:
			capN := int(instr.Operand2)
			fn := int(instr.Operand1)
			caps := vcopy(popN(capN))
			closure := &Closure{
				Fn:       &m.program.fns[fn],
				Captures: caps,
			}
			push(closure)

		case OpPushInt:
			push(Int(instr.Operand1))

		case OpPushGlobal:
			push(m.globals[int(instr.Operand1)])

		case OpPushGlobalRef:
			push(ValueRef{&m.globals[int(instr.Operand1)]})

		case OpPushLocal:
			push(f.Locals[int(instr.Operand1)])

		case OpPushLocalRef:
			push(ValueRef{&f.Locals[int(instr.Operand1)]})

		case OpPushArg:
			push(f.Args[int(instr.Operand1)])

		case OpPushArgRef:
			push(ValueRef{&f.Args[int(instr.Operand1)]})

		case OpPushFn:
			push(&m.program.fns[int(instr.Operand1)])

		case OpPushExternFn:
			push(m.program.extFns[int(instr.Operand1)])

		case OpPushLiteral:
			push(m.program.literals[int(instr.Operand1)])

		case OpBinOp:
			operand2 := pop()
			operand1 := pop()
			op := BinOp(instr.Operand1)
			res, err := EvalBinOp(op, operand1, operand2)
			if err != nil {
				return err
			}
			push(res)

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
	case BinLT:
		return Bool(operand1 < operand2), nil
	case BinLE:
		return Bool(operand1 <= operand2), nil
	case BinGT:
		return Bool(operand1 > operand2), nil
	case BinGE:
		return Bool(operand1 >= operand2), nil
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
	case BinLT:
		return Bool(operand1 < operand2), nil
	case BinLE:
		return Bool(operand1 <= operand2), nil
	case BinGT:
		return Bool(operand1 > operand2), nil
	case BinGE:
		return Bool(operand1 >= operand2), nil
	default:
		panic("invalid BinOp")
	}
}
