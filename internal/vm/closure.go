package vm

import "fmt"

type Closure struct {
	fn   *Fn
	caps []ValueRef
}

func (c *Closure) String() string { return "<func>" }
func (c *Closure) Type() string   { return "Func" }

func (c *Closure) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}

func (c *Closure) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Closure calls are handled directly by the VM.
	panic("not called")
}
