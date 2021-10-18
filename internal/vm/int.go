package vm

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

func (i Int) EvalOp(op Op, operand Value) (Value, error) {
	if op == OpUMinus {
		return NewInt(-i.v), nil
	}

	operandInt, ok := operand.(Int)
	if !ok {
		if operandFloat, ok := operand.(Float); ok {
			return NewFloat(float64(i.v)).EvalOp(op, operandFloat)
		}
		return nil, fmt.Errorf(
			"invalid operation between int and %v", TypeName(operand))
	}

	switch op {
	case OpAdd:
		return NewInt(i.v + operandInt.Int64()), nil
	case OpSub:
		return NewInt(i.v - operandInt.Int64()), nil
	case OpMult:
		return NewInt(i.v * operandInt.Int64()), nil
	case OpDiv:
		if operandInt.Int64() == 0 {
			return nil, ErrDivideByZero
		}
		return NewInt(i.v / operandInt.Int64()), nil
	case OpMod:
		if operandInt.Int64() == 0 {
			return nil, ErrDivideByZero
		}
		return NewInt(i.v % operandInt.Int64()), nil
	case OpLT:
		return NewBool(i.v < operandInt.Int64()), nil
	case OpLE:
		return NewBool(i.v <= operandInt.Int64()), nil
	case OpGT:
		return NewBool(i.v > operandInt.Int64()), nil
	case OpGE:
		return NewBool(i.v >= operandInt.Int64()), nil
	case OpEq:
		return NewBool(i.v == operandInt.Int64()), nil
	default:
		return nil, fmt.Errorf("operator not supported by int")
	}
}
