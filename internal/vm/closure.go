package vm

import "fmt"

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

func (c *Closure) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}

func (c *Closure) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}
