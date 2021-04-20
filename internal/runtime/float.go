package runtime

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

func (f Float) EvalBinOp(op BinOp, operand Value) (Value, error) {
	operandFloat, ok := operand.(Float)
	if !ok {
		if operandInt, ok := operand.(Int); ok {
			operandFloat = NewFloat(float64(operandInt.Int64()))
		} else {
			return nil, fmt.Errorf(
				"invalid operation between float and %v", TypeName(operand))
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
