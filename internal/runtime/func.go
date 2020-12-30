package runtime

import "context"

type ExternFn func(
	ctx context.Context, args []Value) ([]Value, error)

type Emitter struct {
	fns     []Fn
	extFns  []ExternFn
	fnStack []*Fn
}

type Fn struct {
	instrs []Instr
}

func NewEmitter() *Emitter {
	return &Emitter{}
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

func (e *Emitter) ToProgram() *Program {
	return &Program{
		fns:    e.fns,
		extFns: e.extFns,
	}
}
