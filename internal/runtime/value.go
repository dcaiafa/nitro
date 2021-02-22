package runtime

import (
	"context"
	"fmt"
	"strconv"
)

type Value interface {
	fmt.Stringer

	isValue()
}

type String string

func (s String) isValue()       {}
func (s String) String() string { return string(s) }

type Int int64

func (i Int) isValue()       {}
func (i Int) String() string { return strconv.FormatInt(int64(i), 10) }

type Float float64

func (f Float) isValue()       {}
func (f Float) String() string { return strconv.FormatFloat(float64(f), 'g', -1, 64) }

type Bool bool

func (b Bool) isValue()       {}
func (b Bool) String() string { return fmt.Sprint(bool(b)) }

type ValueRef struct {
	Ref *Value
}

func (r ValueRef) isValue()       {}
func (r ValueRef) String() string { return "&" + (*r.Ref).String() }

type Closure struct {
	Fn       *Fn
	Captures []ValueRef
}

func (c *Closure) isValue()       {}
func (c *Closure) String() string { return "<func>" }

type ExternFn func(
	ctx context.Context,
	args []Value,
) ([]Value, error)

func (f ExternFn) isValue()       {}
func (f ExternFn) String() string { return "<func>" }

type Fn struct {
	instrs []Instr
}

func (f *Fn) isValue()       {}
func (f *Fn) String() string { return "<func>" }
