package vm

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

func (a *Array) Type() string { return "list" }

func (a *Array) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("list does not support this operation")
}

func (a *Array) Add(v Value) {
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
	switch key := key.(type) {
	case String:
		switch key.String() {
		case "add":
			return NewNativeFn(a.add), nil
		case "add_iter":
			return NewNativeFn(a.add_iter), nil
		case "len":
			return NewNativeFn(a.len), nil
		case "find":
			return NewNativeFn(a.find), nil
		default:
			return nil, fmt.Errorf("list does not have method %q", key.String())
		}

	case Int:
		idx := int(key.Int64())
		if idx < 0 {
			idx = len(a.array) + idx
		}
		if idx < 0 || idx >= len(a.array) {
			return nil, nil
		}
		return a.array[idx], nil

	default:
		return nil, fmt.Errorf(
			"cannot index string using key type %v",
			TypeName(key))
	}
}

func (a *Array) IndexRef(key Value) (ValueRef, error) {
	index, ok := key.(Int)
	if !ok {
		return ValueRef{}, fmt.Errorf(
			"cannot index list: index must be Int, but it is %v",
			TypeName(key))
	}
	if index.Int64() < 0 || index.Int64() > math.MaxInt32 {
		return ValueRef{}, fmt.Errorf(
			"cannot index list: invalid index %v",
			index.Int64())
	}

	i := int(index.Int64())
	if i >= len(a.array) {
		return ValueRef{}, fmt.Errorf(
			"cannot index list: index %v is greater than array size %v",
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

func (a *Array) MakeIterator() Iterator {
	i := &arrayIter{
		arr:  a,
		next: 0,
	}
	return NewIterator(i.Next, nil, 2)
}

type arrayIter struct {
	arr  *Array
	next int
}

func (i *arrayIter) Next(m *VM, args []Value, nret int) ([]Value, error) {
	if i.next >= i.arr.Len() {
		return nil, nil
	}

	idx := i.next
	i.next++

	v := i.arr.Get(idx)

	return []Value{v, NewInt(int64(idx))}, nil
}

var errListAddUsage error = NewInvalidUsageError("<list>.add(any+)")

func (a *Array) add(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) == 0 {
		return nil, errListAddUsage
	}

	for _, arg := range args {
		a.Add(arg)
	}

	if nRet == 1 {
		return []Value{a}, nil
	} else {
		return nil, nil
	}
}

var errListAddIterUsage error = NewInvalidUsageError("<list>.additer(iter)")

func (a *Array) add_iter(m *VM, args []Value, nret int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errListAddIterUsage
	}

	switch arg := args[0].(type) {
	case *Array:
		a.array = append(a.array, arg.array...)
	case Iterator, Iterable:
		iter, err := MakeIterator(m, arg)
		if err != nil {
			return nil, err
		}
		for {
			v, err := m.IterNext(iter, 1)
			if err != nil {
				return nil, err
			}
			if v == nil {
				break
			}
			a.array = append(a.array, v[0])
		}
	}

	if nret == 1 {
		return []Value{a}, nil
	} else {
		return nil, nil
	}
}

var errListLenUsage error = NewInvalidUsageError("<list>.len()")

func (a *Array) len(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 0 {
		return nil, errListLenUsage
	}

	res := len(a.array)

	return []Value{NewInt(int64(res))}, nil
}

var errListContainsUsage error = NewInvalidUsageError("<list>.contains(any)")

func (a *Array) find(m *VM, args []Value, nRet int) ([]Value, error) {
	if len(args) != 1 {
		return nil, errListContainsUsage
	}
	v := args[0]
	for i, entry := range a.array {
		if entry == v {
			return []Value{NewInt(int64(i))}, nil
		}
	}
	return []Value{nil}, nil
}
