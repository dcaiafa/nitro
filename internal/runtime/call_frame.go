package runtime

type tryCatch struct {
	CatchAddr int
}

type frame struct {
	ExpRetN    int
	Fn         *Fn
	Instrs     []Instr
	Args       []Value
	Captures   []ValueRef
	Locals     []Value
	Stack      []Value
	TryCatches []tryCatch
	IP         int
}

func newFrame() *frame {
	return &frame{}
}

func (f *frame) Init(localN int) {
	f.Locals = make([]Value, localN)
}
