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

func (s String) EvalOp(op Op, operand Value) (Value, error) {
	operandStr, ok := operand.(String)
	if !ok {
		return nil, fmt.Errorf(
			"invalid operation between string and %v",
			TypeName(operand))
	}

	switch op {
	case OpAdd:
		return NewString(s.v + operandStr.String()), nil
	case OpLT:
		return NewBool(s.v < operandStr.String()), nil
	case OpLE:
		return NewBool(s.v <= operandStr.String()), nil
	case OpGT:
		return NewBool(s.v > operandStr.String()), nil
	case OpGE:
		return NewBool(s.v >= operandStr.String()), nil
	case OpEq:
		return NewBool(s.v == operandStr.String()), nil
	default:
		return nil, fmt.Errorf("operator not supported by string")
	}
}

func (s String) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("operator not supported by string")
}

func (s String) Iterate() *Iterator {
	i := &stringIter{
		str:  s.v,
		next: 0,
	}
	return NewIterator(i.Next, 2)
}

type stringIter struct {
	str  string
	next int
}

func (i *stringIter) Next(m *VM, args []Value, nret int) ([]Value, error) {
	if i.next >= len(i.str) {
		return []Value{NewBool(false), nil, nil}, nil
	}

	idx := i.next
	i.next++

	v := NewInt(int64(i.str[idx]))

	return []Value{NewBool(true), v, NewInt(int64(idx))}, nil
}
