package runtime

type Array struct {
	array []Value
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) ValueType() string { return "Array" }

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
