package runtime

type tryCatch struct {
	CatchAddr int
}

type callFrame struct {
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

func newCallFrame() *callFrame {
	return &callFrame{}
}

func (f *callFrame) Init(localN int) {
	f.Locals = make([]Value, localN)
}
