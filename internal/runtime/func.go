package runtime

import "context"

type ExternalFunc func(
	ctx context.Context, args []Value) ([]Value, error)

type Emitter struct {
	fns     []Fn
	extFns  []ExternalFunc
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

func (e *Emitter) Emit(op OpCode, operand uint64) {
	instr := MakeInstr(op, operand)
	curFn := e.curFn()
	curFn.instrs = append(curFn.instrs, instr)
}

func (e *Emitter) AddExternalFunc(fn ExternalFunc) int {
	e.extFns = append(e.extFns, fn)
	return len(e.extFns) - 1
}

func (e *Emitter) ToProgram() *Program {
	return &Program{
		fns:    e.fns,
		extFns: e.extFns,
	}
}
