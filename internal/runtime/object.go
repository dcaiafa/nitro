package runtime

import (
	"reflect"
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

func NewObject() *Object {
	return &Object{
		data: make(map[Value]*objectNode),
		list: newObjectNodeList(),
	}
}

func (o *Object) isValue() {}

func (o *Object) Len() int {
	return len(o.data)
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

func (o *Object) Get(k Value) Value {
	n := o.data[k]
	if n == nil {
		return nil
	}
	return n.value
}

func (o *Object) GetRef(k Value) *Value {
	n := o.data[k]
	if n == nil {
		n = &objectNode{key: k}
		n.InsertAfter(o.list.prev)
		o.data[k] = n
	}
	return &n.value
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
		of.str(`"` + string(v) + `"`)
	case Int:
		of.str(strconv.FormatInt(int64(v), 10))
	case Float:
		of.str(strconv.FormatFloat(float64(v), 'g', -1, 64))
	case Bool:
		if bool(v) {
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
				of.str(string(ks))
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
			of.str(v.Get(i).String())
		}
		of.str("]")

	default:
		of.str("<" + reflect.TypeOf(v).String() + ">")
	}
}

func (of *objectFormatter) str(s string) {
	of.w.WriteString(s)
}
