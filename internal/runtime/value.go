package runtime

import (
	"context"
	"fmt"
	"strconv"
)

type Value interface {
	fmt.Stringer

	ValueType() string
}

type String string

func (s String) ValueType() string { return "String" }
func (s String) String() string    { return string(s) }

type Int int64

func (i Int) ValueType() string { return "Int" }
func (i Int) String() string    { return strconv.FormatInt(int64(i), 10) }

type Float float64

func (f Float) ValueType() string { return "Float" }
func (f Float) String() string    { return strconv.FormatFloat(float64(f), 'g', -1, 64) }

type Bool bool

func (b Bool) ValueType() string { return "Bool" }
func (b Bool) String() string    { return fmt.Sprint(bool(b)) }

type ValueRef struct {
	Ref *Value
}

func (r ValueRef) ValueType() string { return "&" + (*r.Ref).ValueType() }
func (r ValueRef) String() string    { return "&" + (*r.Ref).String() }

type Closure struct {
	Fn       *Fn
	Captures []ValueRef
}

func (c *Closure) ValueType() string { return "Func" }
func (c *Closure) String() string    { return "<func>" }

type ExternFn func(
	ctx context.Context,
	args []Value,
) ([]Value, error)

func (f ExternFn) ValueType() string { return "Func" }
func (f ExternFn) String() string    { return "<func>" }

type Fn struct {
	instrs []Instr
}

func (f *Fn) ValueType() string { return "Func" }
func (f *Fn) String() string    { return "<func>" }
