package vm

type Closure struct {
	fn   *Fn
	caps []ValueRef
}

func (c *Closure) String() string { return "<func>" }
func (c *Closure) Type() string   { return "Func" }
func (c *Closure) Traits() Traits { return TraitNone }

func (c *Closure) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Closure calls are handled directly by the VM.
	panic("not called")
}
