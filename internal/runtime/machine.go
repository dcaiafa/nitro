package runtime

import (
	"context"
	"fmt"
	"reflect"
)

type OpCode byte

const (
	OpNop OpCode = iota
	OpCall
	OpMakeClosure
	OpPushInt
	OpPushLocal
	OpPushLocalRef
	OpPushArg
	OpPushArgRef
	OpPushExternFn
	OpPushLiteral
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
}

func NewMachine() *Machine {
	return &Machine{}
}

func (m *Machine) Run(ctx context.Context, p *Program) error {
	m.program = p
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
			push(int64(instr.Operand1))

		case OpPushLocal:
			push(f.Locals[int(instr.Operand1)])

		case OpPushLocalRef:
			push(&f.Locals[int(instr.Operand1)])

		case OpPushArg:
			push(f.Args[int(instr.Operand1)])

		case OpPushArgRef:
			push(&f.Args[int(instr.Operand1)])

		case OpPushExternFn:
			push(m.program.extFns[int(instr.Operand1)])

		case OpPushLiteral:
			push(m.program.literals[int(instr.Operand1)])

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
			rval := pop().(*Value)
			*rval = val

		case OpInitCallFrame:
			localN := instr.Operand1
			f.Init(int(localN))

		default:
			panic("invalid instruction")
		}

		ip++
	}
}
