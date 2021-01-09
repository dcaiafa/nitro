package runtime

type Emitter struct {
	fnStack   []*Fn
	stringMap map[string]int

	fns      []Fn
	extFns   []ExternFn
	literals []Value
}

func NewEmitter() *Emitter {
	return &Emitter{
		stringMap: make(map[string]int),
	}
}

func (e *Emitter) NewFn() int {
	e.fns = append(e.fns, Fn{})
	return len(e.fns) - 1
}

func (e *Emitter) PushFn(fn int) {
	e.fnStack = append(e.fnStack, &e.fns[fn])
}

func (e *Emitter) PopFn() {
	e.fnStack = e.fnStack[:len(e.fnStack)-1]
}

func (e *Emitter) curFn() *Fn {
	return e.fnStack[len(e.fnStack)-1]
}

func (e *Emitter) Emit(op OpCode, operand1 uint16, operand2 byte) {
	curFn := e.curFn()
	curFn.instrs = append(curFn.instrs, Instr{
		Op:       op,
		Operand2: operand2,
		Operand1: operand1,
	})
}

func (e *Emitter) AddExternalFunc(fn ExternFn) int {
	e.extFns = append(e.extFns, fn)
	return len(e.extFns) - 1
}

func (e *Emitter) AddString(s string) int {
	n, ok := e.stringMap[s]
	if !ok {
		n = e.AddLiteral(s)
		e.stringMap[s] = n
	}
	return n
}

func (e *Emitter) AddLiteral(v Value) int {
	e.literals = append(e.literals, v)
	return len(e.literals) - 1
}

func (e *Emitter) ToProgram() *Program {
	return &Program{
		fns:      e.fns,
		extFns:   e.extFns,
		literals: e.literals,
	}
}
