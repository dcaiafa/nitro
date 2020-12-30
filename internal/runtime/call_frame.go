package runtime

type CallFrameFactory struct {
}

func (f *CallFrameFactory) NewCallFrame() *CallFrame {
	return &CallFrame{}
}

type CallFrame struct {
	Closure *Closure
	ExpRetN int
	Args    []Value
	Locals  []Value
	Stack   []Value
	IP      int
}

func (f *CallFrame) Init(localN int) {
	f.Locals = make([]Value, localN)
}
