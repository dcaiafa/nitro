package runtime

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

func (o *Object) ForEach(f func(k, v Value) bool) {
	for n := o.list.next; n != o.list; n = n.next {
		if !f(n.key, n.value) {
			return
		}
	}
}
