package runtime

import (
	"context"
	"fmt"
	"strconv"
)

type Kind int

const (
	UndefinedKind Kind = iota
	NilKind
	StringKind
	IntKind
	FloatKind
	BoolKind
	RefKind
	FuncKind
	ArrayKind
	ObjectKind
)

type Value interface {
	fmt.Stringer

	Type() string
	Kind() Kind
}

type String string

func (s String) String() string { return string(s) }
func (s String) Type() string   { return "String" }
func (s String) Kind() Kind     { return StringKind }

type Int int64

func (i Int) String() string { return strconv.FormatInt(int64(i), 10) }
func (i Int) Type() string   { return "Int" }
func (i Int) Kind() Kind     { return IntKind }

type Float float64

func (f Float) String() string { return strconv.FormatFloat(float64(f), 'g', -1, 64) }
func (f Float) Type() string   { return "Float" }
func (f Float) Kind() Kind     { return FloatKind }

type Bool bool

func (b Bool) String() string { return fmt.Sprint(bool(b)) }
func (b Bool) Type() string   { return "Bool" }
func (b Bool) Kind() Kind     { return BoolKind }

type ValueRef struct {
	Ref *Value
}

func (r ValueRef) String() string { return "&" + (*r.Ref).String() }
func (r ValueRef) Type() string   { return "&" + (*r.Ref).Type() }
func (r ValueRef) Kind() Kind     { return RefKind }

type Closure struct {
	Fn       *Fn
	ExternFn ExternFn
	Captures []ValueRef
}

func (c *Closure) String() string { return "<func>" }
func (c *Closure) Type() string   { return "Func" }
func (c *Closure) Kind() Kind     { return FuncKind }

type ExternFn func(
	ctx context.Context,
	caps []ValueRef,
	args []Value,
) ([]Value, error)

func (f ExternFn) String() string { return "<func>" }
func (f ExternFn) Type() string   { return "Func" }
func (f ExternFn) Kind() Kind     { return FuncKind }

type Fn struct {
	instrs []Instr
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }
func (f *Fn) Kind() Kind     { return FuncKind }
