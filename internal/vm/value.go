package vm

import (
	"errors"
	"fmt"
	"io"
)

var ErrOperationNotSupported = errors.New("operation not supported")

type Traits uint32

const (
	TraitNone Traits = 0
	TraitEq   Traits = 1 << iota
)

type Value interface {
	String() string
	Type() string
	Traits() Traits
}

type Operable interface {
	Value
	EvalOp(op Op, right Value) (Value, error)
}

type FallbackEvaluator interface {
	Operable
	FallbackEvalOp(op Op, left Value) (Value, error)
}

type Closer interface {
	Value
	io.Closer
}

func TypeName(v Value) string {
	if v == nil {
		return "nil"
	}
	return v.Type()
}

type Callable interface {
	Value
	Call(m *VM, args []Value, nRet int) ([]Value, error)
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
func (r ValueRef) Traits() Traits { return TraitNone }
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

func EvalOp(op Op, operand1, operand2 Value) (Value, error) {
	if operand1 != nil &&
		(operand2 != nil || op == OpUMinus) &&
		(operand1.Traits()&TraitEq != 0) {
		opWasNE := false
		if op == OpNE {
			opWasNE = true
			op = OpEq
		}
		res, err := operand1.(Operable).EvalOp(op, operand2)
		if err != nil {
			if errors.Is(err, ErrOperationNotSupported) {
				if fallback, ok := operand2.(FallbackEvaluator); ok {
					res, err = fallback.EvalOp(op, operand1)
				}
			}
			if err != nil {
				return nil, err
			}
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
