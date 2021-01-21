package runtime

// TODO: reuse callframes, consider pre-allocate Locals, Args, Stack

type CallFrameFactory struct {
}

func (f *CallFrameFactory) NewCallFrame() *CallFrame {
	return &CallFrame{}
}

type CallFrame struct {
	ExpRetN  int
	Instrs   []Instr
	Args     []Value
	Captures []ValueRef
	Locals   []Value
	Stack    []Value
	IP       int
}

func (f *CallFrame) Init(localN int) {
	f.Locals = make([]Value, localN)
}
