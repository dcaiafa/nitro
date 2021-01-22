package runtime

type Object struct {
	data map[Value]Value
}

func NewObject() *Object {
	return &Object{
		data: make(map[Value]Value),
	}
}

func (o *Object) isValue() {}

func (o *Object) Put(k, v Value) {
	o.data[k] = v
}

func (o *Object) Get(k Value) Value {
	return o.data[k]
}
