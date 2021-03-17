package runtime

import (
	"context"
	"fmt"
	"math"
)

type Array struct {
	array []Value
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) Type() string { return "Array" }

func (a *Array) Append(v Value) {
	a.array = append(a.array, v)
}

func (a *Array) Get(index int) Value {
	if index >= len(a.array) {
		return nil
	}
	return a.array[index]
}

func (a *Array) Index(key Value) (Value, error) {
	index, ok := key.(Int)
	if !ok {
		return nil, fmt.Errorf(
			"cannot index array: index must be Int, but it is %v",
			key.Type())
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
			key.Type())
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

func (a *Array) Len() int {
	return len(a.array)
}

func (a *Array) String() string {
	return formatObject(a)
}

func (a *Array) Enumerate() *Closure {
	var arr Value = a
	var next Value = NewInt(0)
	return NewClosure(
		arrayIter,
		[]ValueRef{NewValueRef(&arr), NewValueRef(&next)})
}

func arrayIter(ctx context.Context, caps []ValueRef, args []Value, expRetN int) ([]Value, error) {
	var (
		arr  = (*caps[0].Ref).(*Array)
		next = (*caps[1].Ref).(Int)
	)

	if int(next.Int64()) >= arr.Len() {
		return []Value{nil, NewBool(false)}, nil
	}

	*caps[1].Ref = NewInt(int64(next.Int64() + 1))

	v, err := arr.Index(next)
	if err != nil {
		return nil, err
	}

	return []Value{v, NewBool(true)}, nil
}
