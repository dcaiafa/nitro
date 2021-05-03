package vm

import (
	"fmt"
)

type Value interface {
	fmt.Stringer

	Type() string
	EvalOp(op Op, operand Value) (Value, error)
}

func TypeName(v Value) string {
	if v == nil {
		return "nil"
	}
	return v.Type()
}

type Callable interface {
	Value
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

func (r ValueRef) EvalOp(op Op, operand Value) (Value, error) {
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

func EvalOp(op Op, operand1, operand2 Value) (Value, error) {
	if operand1 != nil && (operand2 != nil || op == OpUMinus) {
		opWasNE := false
		if op == OpNE {
			opWasNE = true
			op = OpEq
		}
		res, err := operand1.EvalOp(op, operand2)
		if err != nil {
			return nil, err
		}
		if opWasNE {
			resb, ok := res.(Bool)
			if !ok {
				return nil, fmt.Errorf(
					"evaluation of Eq resulted in %q instead of bool",
					TypeName(res))
			}
			res = NewBool(!resb.Bool())
		}
		return res, nil
	}

	switch op {
	case OpEq:
		return NewBool(operand1 == operand2), nil
	case OpNE:
		return NewBool(operand1 != operand2), nil
	default:
		return nil, fmt.Errorf(
			"operation not supported between %v and %v",
			TypeName(operand1), TypeName(operand2))
	}
}
