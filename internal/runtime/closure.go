package runtime

type Closure struct {
	fn    *Fn
	extFn NativeFn
	caps  []ValueRef
}

func NewClosure(extFn NativeFn, caps []ValueRef) *Closure {
	return &Closure{
		extFn: extFn,
		caps:  caps,
	}
}

func (c *Closure) String() string { return "<func>" }
func (c *Closure) Type() string   { return "Func" }
func (c *Closure) isCallable()    {}
