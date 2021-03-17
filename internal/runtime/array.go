package runtime

import "context"

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

func (a *Array) GetRef(index int) *Value {
	if index >= len(a.array) {
		return nil
	}
	return &a.array[index]
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
		next = int((*caps[1].Ref).(Int).Int64())
	)

	if next >= arr.Len() {
		return []Value{nil, NewBool(false)}, nil
	}

	*caps[1].Ref = NewInt(int64(next + 1))

	return []Value{arr.Get(next), NewBool(true)}, nil
}
