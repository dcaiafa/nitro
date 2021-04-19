package runtime

type NativeFn func(
	m *Machine,
	caps []ValueRef,
	args []Value,
	retN int,
) ([]Value, error)

func (f NativeFn) String() string { return "<func>" }
func (f NativeFn) Type() string   { return "Func" }
func (f NativeFn) isCallable()    {}
