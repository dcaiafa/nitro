package runtime

import (
	"fmt"
	"strconv"
)

type Value interface {
	fmt.Stringer

	Type() string
}

func TypeName(v Value) string {
	if v == nil {
		return "nil"
	}
	return v.Type()
}

type Callable interface {
	Value

	isCallable()
}

type Indexable interface {
	Value
	Index(key Value) (Value, error)
	IndexRef(key Value) (ValueRef, error)
}

type Evaluator interface {
	Value
	EvalBinOp(op BinOp, operand Value) (Value, error)
	EvalUnaryMinus() (Value, error)
}

type nilValue struct{}

func (v nilValue) String() string { return "<nil>" }
func (v nilValue) Type() string   { return "nil" }

func wrapNil(v Value) Value {
	if v == nil {
		return nilValue{}
	}
	return v
}

type String struct {
	v string
}

func NewString(v string) String { return String{v} }

func (s String) String() string { return s.v }
func (s String) Type() string   { return "String" }
func (s String) Len() int       { return len(s.v) }

func (s String) Index(key Value) (Value, error) {
	idxValue, ok := key.(Int)
	if !ok {
		return nil, fmt.Errorf(
			"cannot index string: index must be Int, but it is %v",
			key.Type())
	}

	idx := int(idxValue.Int64())
	if idx < 0 {
		idx = len(s.v) + idx
	}
	if idx < 0 || idx >= len(s.v) {
		return nil, nil
	}
	return NewInt(int64(s.v[idx])), nil
}

func (s String) IndexRef(key Value) (ValueRef, error) {
	return NewValueRef(nil), fmt.Errorf("cannot modify string")
}

func (s String) Slice(b, e Value) (Value, error) {
	bi, ok := b.(Int)
	ei, ok2 := e.(Int)
	if !ok || !ok2 {
		return nil, fmt.Errorf(
			"slice indices must be Int; instead they are %q and %q",
			TypeName(b), TypeName(e))
	}

	begin := int(bi.Int64())
	end := int(ei.Int64())

	if begin < 0 {
		return nil, fmt.Errorf(
			"invalid slice begin index %v; begin index must be >= 0",
			begin)
	}

	if end < 0 {
		end = len(s.v) + end
	}
	if end > len(s.v) {
		end = len(s.v)
	}
	if end < begin {
		begin = 0
		end = 0
	}

	return NewString(s.v[begin:end]), nil
}

func (s String) EvalBinOp(op BinOp, operand Value) (Value, error) {
	if op == BinEq {
		return NewBool(s == operand), nil
	} else if op == BinNE {
		return NewBool(s != operand), nil
	}

	operandStr, ok := operand.(String)
	if !ok {
		return nil, fmt.Errorf(
			"invalid operation between string and %v",
			TypeName(operand))
	}

	switch op {
	case BinAdd:
		return NewString(s.v + operandStr.String()), nil
	case BinLT:
		return NewBool(s.v < operandStr.String()), nil
	case BinLE:
		return NewBool(s.v <= operandStr.String()), nil
	case BinGT:
		return NewBool(s.v > operandStr.String()), nil
	case BinGE:
		return NewBool(s.v >= operandStr.String()), nil
	case BinEq:
		return NewBool(s.v == operandStr.String()), nil
	case BinNE:
		return NewBool(s.v != operandStr.String()), nil
	default:
		return nil, fmt.Errorf("operator not supported by string")
	}
}

func (s String) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("operator not supported by string")
}

type Int struct {
	v int64
}

func NewInt(v int64) Int { return Int{v} }

func (i Int) Int64() int64   { return i.v }
func (i Int) String() string { return strconv.FormatInt(i.v, 10) }
func (i Int) Type() string   { return "Int" }

func (i Int) EvalBinOp(op BinOp, operand Value) (Value, error) {
	operandInt, ok := operand.(Int)
	if !ok {
		if operandFloat, ok := operand.(Float); ok {
			return NewFloat(float64(i.v)).EvalBinOp(op, operandFloat)
		}
	}

	if op == BinEq {
		return NewBool(i == operandInt), nil
	} else if op == BinNE {
		return NewBool(i != operandInt), nil
	}

	switch op {
	case BinAdd:
		return NewInt(i.v + operandInt.Int64()), nil
	case BinSub:
		return NewInt(i.v - operandInt.Int64()), nil
	case BinMult:
		return NewInt(i.v * operandInt.Int64()), nil
	case BinDiv:
		return NewInt(i.v / operandInt.Int64()), nil
	case BinMod:
		return NewInt(i.v % operandInt.Int64()), nil
	case BinLT:
		return NewBool(i.v < operandInt.Int64()), nil
	case BinLE:
		return NewBool(i.v <= operandInt.Int64()), nil
	case BinGT:
		return NewBool(i.v > operandInt.Int64()), nil
	case BinGE:
		return NewBool(i.v >= operandInt.Int64()), nil
	case BinEq:
		return NewBool(i.v == operandInt.Int64()), nil
	case BinNE:
		return NewBool(i.v != operandInt.Int64()), nil
	default:
		return nil, fmt.Errorf("operator not supported by int")
	}
}

func (i Int) EvalUnaryMinus() (Value, error) {
	return NewInt(-i.v), nil
}

type Float struct {
	v float64
}

func NewFloat(v float64) Float { return Float{v} }

func (f Float) Float64() float64 { return f.v }
func (f Float) String() string   { return strconv.FormatFloat(f.v, 'g', -1, 64) }
func (f Float) Type() string     { return "Float" }

func (f Float) EvalBinOp(op BinOp, operand Value) (Value, error) {
	operandFloat, ok := operand.(Float)
	if !ok {
		if operandInt, ok := operand.(Int); ok {
			operandFloat = NewFloat(float64(operandInt.Int64()))
		}
	}

	if op == BinEq {
		return NewBool(f == operandFloat), nil
	} else if op == BinNE {
		return NewBool(f != operandFloat), nil
	}

	switch op {
	case BinAdd:
		return NewFloat(f.v + operandFloat.Float64()), nil
	case BinSub:
		return NewFloat(f.v - operandFloat.Float64()), nil
	case BinMult:
		return NewFloat(f.v * operandFloat.Float64()), nil
	case BinDiv:
		return NewFloat(f.v / operandFloat.Float64()), nil
	case BinLT:
		return NewBool(f.v < operandFloat.Float64()), nil
	case BinLE:
		return NewBool(f.v <= operandFloat.Float64()), nil
	case BinGT:
		return NewBool(f.v > operandFloat.Float64()), nil
	case BinGE:
		return NewBool(f.v >= operandFloat.Float64()), nil
	case BinEq:
		return NewBool(f.v == operandFloat.Float64()), nil
	case BinNE:
		return NewBool(f.v != operandFloat.Float64()), nil
	default:
		return nil, fmt.Errorf("operator not supported by float")
	}
}

func (f Float) EvalUnaryMinus() (Value, error) {
	return NewFloat(-f.v), nil
}

type Bool struct {
	v bool
}

func NewBool(v bool) Bool { return Bool{v} }

func (b Bool) Bool() bool     { return b.v }
func (b Bool) String() string { return fmt.Sprint(b.v) }
func (b Bool) Type() string   { return "Bool" }

type ValueRef struct {
	Ref *Value
}

func NewValueRef(ref *Value) ValueRef {
	return ValueRef{Ref: ref}
}

func (r ValueRef) String() string { return "&" + (*r.Ref).String() }
func (r ValueRef) Type() string   { return "&" + TypeName(*r.Ref) }
func (r ValueRef) Refo() *Value   { return r.Ref }

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

type Iterator struct {
	fn         *Fn
	extFn      NativeFn
	captures   []ValueRef
	iterNRet   int
	locals     []Value
	tryCatches []tryCatch
	defers     []*Closure
	ip         int
}

func (e *Iterator) String() string { return "<Iterator>" }
func (e *Iterator) Type() string   { return "Iterator" }
func (e *Iterator) isCallable()    {}
func (e *Iterator) IterNRet() int  { return e.iterNRet }

func NewIterator(extFn NativeFn, caps []ValueRef, nret int) *Iterator {
	return &Iterator{
		extFn:    extFn,
		captures: caps,
		iterNRet: nret,
	}
}

type NativeFn func(
	m *Machine,
	caps []ValueRef,
	args []Value,
	retN int,
) ([]Value, error)

func (f NativeFn) String() string { return "<func>" }
func (f NativeFn) Type() string   { return "Func" }
func (f NativeFn) isCallable()    {}

type Fn struct {
	name      int
	locations []Location
	instrs    []Instr
	minArgs   int
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }
func (f *Fn) isCallable()    {}

func ToString(v Value) string {
	if v == nil {
		return "<nil>"
	}
	return v.String()
}

func CoerceToBool(v Value) bool {
	switch v := v.(type) {
	case Bool:
		return v.Bool()
	default:
		return v != nil
	}
}

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	evaluator, ok := operand1.(Evaluator)
	if ok {
		return evaluator.EvalBinOp(op, operand2)
	} else if op == BinEq {
		return NewBool(operand1 == operand2), nil
	} else if op == BinNE {
		return NewBool(operand1 != operand2), nil
	} else {
		return nil, fmt.Errorf(
			"left value with type %q does not support binary operation",
			TypeName(operand1))
	}
}
