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

type String struct {
	v string
}

func NewString(v string) String { return String{v} }

func (s String) String() string { return s.v }
func (s String) Type() string   { return "String" }
func (s String) Kind() Kind     { return StringKind }

type Int struct {
	v int64
}

func NewInt(v int64) Int { return Int{v} }

func (i Int) Int64() int64   { return i.v }
func (i Int) String() string { return strconv.FormatInt(i.v, 10) }
func (i Int) Type() string   { return "Int" }
func (i Int) Kind() Kind     { return IntKind }

type Float struct {
	v float64
}

func NewFloat(v float64) Float { return Float{v} }

func (f Float) Float64() float64 { return f.v }
func (f Float) String() string   { return strconv.FormatFloat(f.v, 'g', -1, 64) }
func (f Float) Type() string     { return "Float" }
func (f Float) Kind() Kind       { return FloatKind }

type Bool struct {
	v bool
}

func NewBool(v bool) Bool { return Bool{v} }

func (b Bool) Bool() bool     { return b.v }
func (b Bool) String() string { return fmt.Sprint(b.v) }
func (b Bool) Type() string   { return "Bool" }
func (b Bool) Kind() Kind     { return BoolKind }

type ValueRef struct {
	Ref *Value
}

func NewValueRef(ref *Value) ValueRef {
	return ValueRef{Ref: ref}
}

func (r ValueRef) Refo() *Value   { return r.Ref }
func (r ValueRef) String() string { return "&" + (*r.Ref).String() }
func (r ValueRef) Type() string   { return "&" + (*r.Ref).Type() }
func (r ValueRef) Kind() Kind     { return RefKind }

type Closure struct {
	fn    *Fn
	extFn ExternFn
	caps  []ValueRef
}

func NewClosure(extFn ExternFn, caps []ValueRef) *Closure {
	return &Closure{
		extFn: extFn,
		caps:  caps,
	}
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
	locations []Location
	instrs    []Instr
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }
func (f *Fn) Kind() Kind     { return FuncKind }
