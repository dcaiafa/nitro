package runtime

type Object struct {
	fallback func(key interface{}) (interface{}, bool)
	data     map[interface{}]interface{}
}

func NewObject() *Object {
	return &Object{
		data: make(map[interface{}]interface{}),
	}
}

func (o *Object) Get(key interface{}) (interface{}, bool) {
	v, ok := o.data[key]
	if !ok && o.fallback != nil {
		return o.fallback(key)
	}
	return v, ok
}

func (o *Object) Put(key, value interface{}) {
	o.data[key] = value
}
