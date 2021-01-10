package runtime

import "context"

type Value interface {
	isValue()
}

type String string

func (s String) isValue() {}

type Int int64

func (i Int) isValue() {}

type Float float64

func (f Float) isValue() {}

type Bool bool

func (b Bool) isValue() {}

type ValueRef struct {
	Ref *Value
}

func (r ValueRef) isValue() {}

type Closure struct {
	Fn       *Fn
	Captures []Value
}

func (c *Closure) isValue() {}

type ExternFn func(
	ctx context.Context,
	args []Value,
) ([]Value, error)

func (f ExternFn) isValue() {}

type Fn struct {
	instrs []Instr
}

func (f *Fn) isValue() {}
