package runtime

type Label struct {
	fn   *Fn
	addr int
	refs []int
}

type Emitter struct {
	fnStack   []*Fn
	stringMap map[string]int

	globals  int
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

func (e *Emitter) NewLabel() *Label {
	return &Label{
		fn:   e.curFn(),
		addr: -1,
	}
}

func (e *Emitter) AssignLabel(label *Label) {
	if label.addr != -1 {
		panic("label already attached")
	}
	// The label is placed *after* the last instruction.
	instrs := e.curFn().instrs
	label.addr = len(instrs)
	for _, ref := range label.refs {
		instr := &instrs[ref]
		instr.Operand1, instr.Operand2 = word24ToOperands(uint32(label.addr))
	}
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

func (e *Emitter) EmitJump(op OpCode, label *Label) {
	if label.addr == -1 {
		if e.curFn() != label.fn {
			panic("cannot jump across fns")
		}
		e.Emit(op, 0, 0)
		instrs := e.curFn().instrs
		lastEmittedInstr := len(instrs) - 1
		label.refs = append(label.refs, lastEmittedInstr)
	} else {
		// TODO: Check that fns are no longer than 2^24-1.
		operand1, operand2 := word24ToOperands(uint32(label.addr))
		e.Emit(op, operand1, operand2)
	}
}

func (e *Emitter) AddExternalFunc(fn ExternFn) int {
	e.extFns = append(e.extFns, fn)
	return len(e.extFns) - 1
}

func (e *Emitter) AddString(s string) int {
	n, ok := e.stringMap[s]
	if !ok {
		n = e.AddLiteral(String(s))
		e.stringMap[s] = n
	}
	return n
}

func (e *Emitter) AddLiteral(v Value) int {
	e.literals = append(e.literals, v)
	return len(e.literals) - 1
}

func (e *Emitter) SetGlobalCount(n int) {
	e.globals = n
}

func (e *Emitter) ToProgram() *Program {
	return &Program{
		globals:  e.globals,
		fns:      e.fns,
		extFns:   e.extFns,
		literals: e.literals,
	}
}
