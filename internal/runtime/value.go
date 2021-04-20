package runtime

import (
	"fmt"
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

type ValueRef struct {
	Ref *Value
}

func NewValueRef(ref *Value) ValueRef {
	return ValueRef{Ref: ref}
}

func (r ValueRef) String() string { return "&" + (*r.Ref).String() }
func (r ValueRef) Type() string   { return "&" + TypeName(*r.Ref) }
func (r ValueRef) Refo() *Value   { return r.Ref }

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
		evaluator, ok := operand1.(Evaluator)
		if ok {
			return evaluator.EvalBinOp(op, operand2)
		}
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
