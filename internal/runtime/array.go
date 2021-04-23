package runtime

import (
	"fmt"
	"math"
)

type Array struct {
	array []Value
}

func NewArray() *Array {
	return &Array{}
}

func NewArrayWithSlice(s []Value) *Array {
	return &Array{array: s}
}

func (a *Array) Type() string { return "Array" }

func (a *Array) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("array does not support this operation")
}

func (a *Array) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("array does not support this operation")
}

func (a *Array) Push(v Value) {
	a.array = append(a.array, v)
}

func (a *Array) Back() Value {
	if len(a.array) == 0 {
		return nil
	}
	return a.array[len(a.array)-1]
}

func (a *Array) Get(index int) Value {
	if index >= len(a.array) {
		return nil
	}
	return a.array[index]
}

func (a *Array) Put(index int, v Value) {
	a.array[index] = v
}

func (a *Array) Index(key Value) (Value, error) {
	index, ok := key.(Int)
	if !ok {
		return nil, fmt.Errorf(
			"cannot index array: index must be Int, but it is %v",
			TypeName(key))
	}
	if index.Int64() < 0 || index.Int64() > math.MaxInt32 {
		return nil, fmt.Errorf(
			"cannot index array: invalid index %v",
			index.Int64())
	}

	i := int(index.Int64())
	if i >= len(a.array) {
		return nil, nil
	}
	return a.array[i], nil
}

func (a *Array) IndexRef(key Value) (ValueRef, error) {
	index, ok := key.(Int)
	if !ok {
		return ValueRef{}, fmt.Errorf(
			"cannot index array: index must be Int, but it is %v",
			TypeName(key))
	}
	if index.Int64() < 0 || index.Int64() > math.MaxInt32 {
		return ValueRef{}, fmt.Errorf(
			"cannot index array: invalid index %v",
			index.Int64())
	}

	i := int(index.Int64())
	if i >= len(a.array) {
		return ValueRef{}, fmt.Errorf(
			"cannot index array: index %v is greater than array size %v",
			i, len(a.array))
	}
	return NewValueRef(&a.array[i]), nil
}

func (a *Array) Slice(b, e Value) (Value, error) {
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
		end = len(a.array) + end
	}
	if end > len(a.array) {
		end = len(a.array)
	}
	if end < begin {
		begin = 0
		end = 0
	}

	return NewArrayWithSlice(a.array[begin:end]), nil
}

func (a *Array) Len() int {
	return len(a.array)
}

func (a *Array) String() string {
	return formatObject(a)
}

func (a *Array) Iterate() *Iterator {
	var arr Value = a
	var next Value = NewInt(0)
	return NewIterator(
		arrayIter,
		[]ValueRef{NewValueRef(&arr), NewValueRef(&next)},
		1)
}

func arrayIter(m *Machine, caps []ValueRef, args []Value, nRet int) ([]Value, error) {
	var (
		arr  = (*caps[0].Ref).(*Array)
		next = (*caps[1].Ref).(Int)
	)

	if int(next.Int64()) >= arr.Len() {
		return []Value{NewBool(false), nil}, nil
	}

	*caps[1].Ref = NewInt(int64(next.Int64() + 1))

	v, err := arr.Index(next)
	if err != nil {
		return nil, err
	}

	return []Value{NewBool(true), v}, nil
}
