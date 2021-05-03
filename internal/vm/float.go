package vm

import (
	"fmt"
	"strconv"
)

type Float struct {
	v float64
}

func NewFloat(v float64) Float { return Float{v} }

func (f Float) Float64() float64 { return f.v }
func (f Float) String() string   { return strconv.FormatFloat(f.v, 'g', -1, 64) }
func (f Float) Type() string     { return "Float" }

func (f Float) EvalOp(op Op, operand Value) (Value, error) {
	if op == OpUMinus {
		return NewFloat(-f.v), nil
	}

	operandFloat, ok := operand.(Float)
	if !ok {
		if operandInt, ok := operand.(Int); ok {
			operandFloat = NewFloat(float64(operandInt.Int64()))
		} else {
			return nil, fmt.Errorf(
				"invalid operation between float and %v", TypeName(operand))
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
		return nil, fmt.Errorf("operator not supported by float")
	}
}
