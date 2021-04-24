package vm

import "fmt"

type String struct {
	v string
}

func NewString(v string) String { return String{v} }

func (s String) String() string { return s.v }
func (s String) Type() string   { return "String" }
func (s String) Len() int       { return len(s.v) }

func (s String) Index(key Value) (Value, error) {
	idxValue, ok := key.(Int)
	if !ok {
		return nil, fmt.Errorf(
			"cannot index string: index must be Int, but it is %v",
			key.Type())
	}

	idx := int(idxValue.Int64())
	if idx < 0 {
		idx = len(s.v) + idx
	}
	if idx < 0 || idx >= len(s.v) {
		return nil, nil
	}
	return NewInt(int64(s.v[idx])), nil
}

func (s String) IndexRef(key Value) (ValueRef, error) {
	return NewValueRef(nil), fmt.Errorf("cannot modify string")
}

func (s String) Slice(b, e Value) (Value, error) {
	bi, ok := b.(Int)
	ei, ok2 := e.(Int)
	if !ok || !ok2 {
		return nil, fmt.Errorf(
			"slice indices must be Int; instead they are %q and %q",
			TypeName(b), TypeName(e))
	}

	begin := int(bi.Int64())
	end := int(ei.Int64())

	if begin < 0 {
		return nil, fmt.Errorf(
			"invalid slice begin index %v; begin index must be >= 0",
			begin)
	}

	if end < 0 {
		end = len(s.v) + end
	}
	if end > len(s.v) {
		end = len(s.v)
	}
	if end < begin {
		begin = 0
		end = 0
	}

	return NewString(s.v[begin:end]), nil
}

func (s String) EvalBinOp(op BinOp, operand Value) (Value, error) {
	if op == BinEq {
		return NewBool(s == operand), nil
	} else if op == BinNE {
		return NewBool(s != operand), nil
	}

	operandStr, ok := operand.(String)
	if !ok {
		return nil, fmt.Errorf(
			"invalid operation between string and %v",
			TypeName(operand))
	}

	switch op {
	case BinAdd:
		return NewString(s.v + operandStr.String()), nil
	case BinLT:
		return NewBool(s.v < operandStr.String()), nil
	case BinLE:
		return NewBool(s.v <= operandStr.String()), nil
	case BinGT:
		return NewBool(s.v > operandStr.String()), nil
	case BinGE:
		return NewBool(s.v >= operandStr.String()), nil
	case BinEq:
		return NewBool(s.v == operandStr.String()), nil
	case BinNE:
		return NewBool(s.v != operandStr.String()), nil
	default:
		return nil, fmt.Errorf("operator not supported by string")
	}
}

func (s String) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("operator not supported by string")
}
