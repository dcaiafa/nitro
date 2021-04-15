package runtime

import (
	"log"
	"strconv"
	"strings"
)

type objectNode struct {
	next, prev *objectNode
	key        Value
	value      Value
}

func newObjectNodeList() *objectNode {
	l := &objectNode{}
	l.prev = l
	l.next = l
	return l
}

func (n *objectNode) InsertAfter(o *objectNode) {
	n.prev = o
	n.next = o.next
	o.next.prev = n
	o.next = n
}

func (n *objectNode) Remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}

type Object struct {
	data map[Value]*objectNode
	list *objectNode
}

var _ Iterable = (*Object)(nil)

func NewObject() *Object {
	return NewObjectWithCapacity(0)
}

func NewObjectWithCapacity(c int) *Object {
	return &Object{
		data: make(map[Value]*objectNode, c),
		list: newObjectNodeList(),
	}
}

func (o *Object) Type() string { return "Object" }

func (o *Object) Len() int {
	return len(o.data)
}

func (o *Object) Has(k Value) bool {
	_, ok := o.data[k]
	return ok
}

func (o *Object) Get(k Value) (Value, bool) {
	n, ok := o.data[k]
	if !ok {
		return nil, false
	}
	return n.value, true
}

func (o *Object) Put(k, v Value) {
	n := o.data[k]
	if n == nil {
		n = &objectNode{key: k}
		n.InsertAfter(o.list.prev)
		o.data[k] = n
	}
	n.value = v
}

func (o *Object) Index(k Value) (Value, error) {
	n := o.data[k]
	if n == nil {
		return nil, nil
	}
	return n.value, nil
}

func (o *Object) IndexRef(k Value) (ValueRef, error) {
	n := o.data[k]
	if n == nil {
		n = &objectNode{key: k}
		n.InsertAfter(o.list.prev)
		o.data[k] = n
	}
	return NewValueRef(&n.value), nil
}

func (o *Object) GetFirst() (key Value, val Value) {
	if len(o.data) == 0 {
		return nil, nil
	}
	return o.list.next.key, o.list.next.value
}

func (o *Object) GetNext(key Value) (nextKey Value, nextVal Value) {
	node := o.data[key]
	if node == nil || node.next == o.list {
		return nil, nil
	}
	return node.next.key, node.next.value
}

func (o *Object) Delete(key Value) {
	node := o.data[key]
	if node == nil {
		return
	}
	node.Remove()
	delete(o.data, key)
}

func (o *Object) ForEach(f func(k, v Value) bool) {
	for n := o.list.next; n != o.list; n = n.next {
		if !f(n.key, n.value) {
			return
		}
	}
}

func (o *Object) String() string {
	return formatObject(o)
}

func (o *Object) Iterate() *Iterator {
	var obj Value = o
	nextKey, _ := obj.(*Object).GetFirst()
	return NewIterator(
		objectIter,
		[]ValueRef{NewValueRef(&obj), NewValueRef(&nextKey)},
		2)
}

type objectFormatter struct {
	visited map[Value]bool
	w       *strings.Builder
}

func formatObject(v Value) string {
	of := &objectFormatter{
		visited: make(map[Value]bool),
		w:       &strings.Builder{},
	}
	of.format(v)
	return of.w.String()
}

func (of *objectFormatter) format(v Value) {
	switch v := v.(type) {
	case String:
		// TODO: escape special characters.
		of.str(`"` + v.String() + `"`)
	case Int:
		of.str(strconv.FormatInt(v.Int64(), 10))
	case Float:
		of.str(strconv.FormatFloat(v.Float64(), 'g', -1, 64))
	case Bool:
		if v.Bool() {
			of.str("true")
		} else {
			of.str("false")
		}
	case *Object:
		if of.visited[v] {
			of.str("<cycle>")
			return
		}
		of.visited[v] = true
		of.str("{")
		first := true
		v.ForEach(func(k, v Value) bool {
			if !first {
				of.str(", ")
			} else {
				first = false
			}
			if ks, ok := k.(String); ok {
				of.str(ks.String())
			} else {
				of.str("[")
				of.format(k)
				of.str("]")
			}
			of.str(": ")
			of.format(v)
			return true
		})
		of.str("}")
	case *Array:
		of.str("[")
		first := true
		for i := 0; i < v.Len(); i++ {
			if !first {
				of.str(" ")
			} else {
				first = false
			}
			of.str(ToString(v.Get(i)))
		}
		of.str("]")

	default:
		of.str(ToString(v))
	}
}

func (of *objectFormatter) str(s string) {
	of.w.WriteString(s)
}

func objectIter(m *Machine, caps []ValueRef, args []Value, expRetN int) ([]Value, error) {
	var (
		obj = (*caps[0].Ref).(*Object)
		key = *caps[1].Ref
	)

	if expRetN != 2 && expRetN != 3 {
		log.Fatalf("unexpected return count %v", expRetN)
	}

	val, err := obj.Index(key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		if expRetN == 2 {
			return []Value{NewBool(false), nil}, nil
		} else {
			return []Value{NewBool(false), nil, nil}, nil
		}
	}

	nextKey, _ := obj.GetNext(key)
	*caps[1].Ref = nextKey

	if expRetN == 2 {
		return []Value{NewBool(true), key}, nil
	} else {
		return []Value{NewBool(true), key, val}, nil
	}
}
