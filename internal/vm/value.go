package vm

import (
	"fmt"
)

type Value interface {
	fmt.Stringer

	Type() string
	EvalBinOp(op BinOp, operand Value) (Value, error)
	EvalUnaryMinus() (Value, error)
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

type ValueRef struct {
	Ref *Value
}

func NewValueRef(ref *Value) ValueRef {
	return ValueRef{Ref: ref}
}

func (r ValueRef) String() string { return "&" + (*r.Ref).String() }
func (r ValueRef) Type() string   { return "&" + TypeName(*r.Ref) }
func (r ValueRef) Refo() *Value   { return r.Ref }

func (r ValueRef) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("ValueRef does not support this operation")
}

func (r ValueRef) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("ValueRef does not support this operation")
}

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
	if operand1 != nil && operand2 != nil {
		return operand1.EvalBinOp(op, operand2)
	}

	switch op {
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		return nil, fmt.Errorf(
			"operation not supported between %v and %v",
			TypeName(operand1), TypeName(operand2))
	}
}
