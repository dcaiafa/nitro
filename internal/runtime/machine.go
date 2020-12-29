package runtime

import (
	"context"
	"fmt"
)

type CallFrame struct {
	Closure *Closure
	ExpRetN int
	Args    []Value
	Locals  []Value
	Stack   []Value
	IP      int
}

type Machine struct {
	callStack []*CallFrame
	program   *Program
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
		ip     = 0
		instrs = m.program.fns[0].instrs
		f      = &CallFrame{}
	)

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

	m.callStack = append(m.callStack, f)

	for {
		instr := instrs[ip]
		switch instr.Op {
		case OpNop:

		case OpCall:
			expRetN := int(instr.Operand2)
			argN := int(instr.Operand1)
			args := popN(argN)
			closure := pop().(*Closure)

			if closure.External {
				extFn := m.program.extFns[closure.Fn]
				rets, err := extFn(ctx, args)
				if err != nil {
					return err
				}
				if len(rets) < expRetN {
					return fmt.Errorf("expected at least %v returned values", expRetN)
				}
				f.Stack = append(f.Stack, rets[:expRetN]...)
			} else {
				f.IP = ip
				m.callStack = append(m.callStack, f)

				f = &CallFrame{
					Closure: closure,
					ExpRetN: expRetN,
					Args:    args,
				}

				instrs = m.program.fns[closure.Fn].instrs
				ip = 0
			}

		case OpMakeClosure:
			capN := int(instr.Operand2)
			fn := int(instr.Operand1)
			external := fn&0x8000 != 0
			if external {
				fn &= 0x7FFF
			}
			caps := vcopy(popN(capN))
			closure := &Closure{
				External: external,
				Fn:       int(fn),
				Captures: caps,
			}
			push(closure)

		case OpPushInt:
			push(int64(instr.Operand1))

		case OpPushLocal:
			push(f.Locals[int(instr.Operand1)])

		case OpPushLocalRef:
			push(&f.Locals[int(instr.Operand1)])

		case OpRet:
			if f.ExpRetN > len(f.Stack) {
				return fmt.Errorf("error")
			}
			rets := f.Stack[:f.ExpRetN]
			top := len(m.callStack) - 1
			m.callStack = m.callStack[:top]
			if top == 0 {
				return nil
			}
			f = m.callStack[top]
			f.Stack = append(f.Stack, rets...)
			ip = f.IP

		case OpStore:
			val := pop()
			rval := pop().(*Value)
			*rval = val

		default:
			panic("invalid instruction")
		}
		ip++
	}
}
