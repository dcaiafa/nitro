package runtime

import (
	"fmt"
	"strconv"
)

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
