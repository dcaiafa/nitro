package runtime

import (
	"errors"
	"fmt"
)

func EvalBinOp(op BinOp, operand1, operand2 Value) (Value, error) {
	switch operand1 := operand1.(type) {
	case Int:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpInt(op, operand1.Int64(), operand2.Int64())
		case Float:
			return evalBinOpFloat(op, float64(operand1.Int64()), operand2.Float64())
		}
	case Float:
		switch operand2 := operand2.(type) {
		case Int:
			return evalBinOpFloat(op, operand1.Float64(), float64(operand2.Int64()))
		case Float:
			return evalBinOpFloat(op, operand1.Float64(), operand2.Float64())
		}

	case String:
		if operand2, ok := operand2.(String); ok {
			return evalBinOpString(op, operand1.String(), operand2.String())
		}
	}

	return nil, fmt.Errorf(
		"could not evaluate binary expression: type %v is incompatible with type %v",
		operand1.Type(), operand2.Type())
}

func evalBinOpInt(op BinOp, operand1, operand2 int64) (Value, error) {
	switch op {
	case BinAdd:
		return NewInt(operand1 + operand2), nil
	case BinSub:
		return NewInt(operand1 - operand2), nil
	case BinMult:
		return NewInt(operand1 * operand2), nil
	case BinDiv:
		return NewInt(operand1 / operand2), nil
	case BinMod:
		return NewInt(operand1 % operand2), nil
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func evalBinOpFloat(op BinOp, operand1, operand2 float64) (Value, error) {
	switch op {
	case BinAdd:
		return NewFloat(operand1 + operand2), nil
	case BinSub:
		return NewFloat(operand1 - operand2), nil
	case BinMult:
		return NewFloat(operand1 * operand2), nil
	case BinDiv:
		return NewFloat(operand1 / operand2), nil
	case BinMod:
		return nil, errors.New("modulo operation not permitted with Float")
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		panic("invalid BinOp")
	}
}

func evalBinOpString(op BinOp, operand1, operand2 string) (Value, error) {
	switch op {
	case BinAdd:
		return NewString(operand1 + operand2), nil
	case BinLT:
		return NewBool(operand1 < operand2), nil
	case BinLE:
		return NewBool(operand1 <= operand2), nil
	case BinGT:
		return NewBool(operand1 > operand2), nil
	case BinGE:
		return NewBool(operand1 >= operand2), nil
	case BinEq:
		return NewBool(operand1 == operand2), nil
	case BinNE:
		return NewBool(operand1 != operand2), nil
	default:
		return nil, fmt.Errorf("cannot use this operator with string operands")
	}
}

func coerceToBool(v Value) bool {
	switch v := v.(type) {
	case Bool:
		return v.Bool()
	case Int:
		return v.Int64() != 0
	case Float:
		return v.Float64() != 0
	case String:
		return len(v.String()) != 0
	case *Object:
		return v.Len() != 0
	case *Array:
		return v.Len() != 0
	default:
		return v != nil
	}
}
