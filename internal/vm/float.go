package vm

import (
	"strconv"
)

type Float struct {
	v float64
}

func NewFloat(v float64) Float { return Float{v} }

func (f Float) Float64() float64 { return f.v }
func (f Float) String() string   { return strconv.FormatFloat(f.v, 'g', -1, 64) }
func (f Float) Type() string     { return "float" }
func (f Float) Traits() Traits   { return TraitEq }

func (f Float) EvalOp(op Op, operand Value) (Value, error) {
	if op == OpUMinus {
		return NewFloat(-f.v), nil
	}

	operandFloat, ok := operand.(Float)
	if !ok {
		if operandInt, ok := operand.(Int); ok {
			operandFloat = NewFloat(float64(operandInt.Int64()))
		} else if op == OpEq {
			return NewBool(false), nil
		} else {
			return nil, ErrOperationNotSupported
		}
	}

	switch op {
	case OpAdd:
		return NewFloat(f.v + operandFloat.Float64()), nil
	case OpSub:
		return NewFloat(f.v - operandFloat.Float64()), nil
	case OpMult:
		return NewFloat(f.v * operandFloat.Float64()), nil
	case OpDiv:
		if operandFloat.Float64() == 0 {
			return nil, ErrDivideByZero
		}
		return NewFloat(f.v / operandFloat.Float64()), nil
	case OpLT:
		return NewBool(f.v < operandFloat.Float64()), nil
	case OpLE:
		return NewBool(f.v <= operandFloat.Float64()), nil
	case OpGT:
		return NewBool(f.v > operandFloat.Float64()), nil
	case OpGE:
		return NewBool(f.v >= operandFloat.Float64()), nil
	case OpEq:
		return NewBool(f.v == operandFloat.Float64()), nil
	default:
		return nil, ErrOperationNotSupported
	}
}
