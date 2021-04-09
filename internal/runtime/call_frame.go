package runtime

type tryCatch struct {
	CatchAddr int
}

type frame struct {
	ExpRetN    int
	Iterator *Iterator
	Fn         *Fn
	Instrs     []Instr
	Args       []Value
	Captures   []ValueRef
	Locals     []Value
	Stack      []Value
	TryCatches []tryCatch
	Defers     []*Closure
	IP         int
}

func (f *frame) Init(localN int) {
	f.Locals = make([]Value, localN)
}
