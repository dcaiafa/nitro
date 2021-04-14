package runtime

import (
	"github.com/dcaiafa/nitro/internal/token"
)

type Label struct {
	fn   *Fn
	addr int
	refs []int
}

type Param struct {
	global int
}

type Emitter struct {
	fnStack   []*Fn
	stringMap map[string]int

	globals   int
	fns       []Fn
	extFns    []ExternFn
	literals  []Value
	params    map[string]*Param
	reqParamN int
	curFile   string
	curFileID int
}

func NewEmitter() *Emitter {
	return &Emitter{
		stringMap: make(map[string]int),
		params:    make(map[string]*Param),
	}
}

func (e *Emitter) AddGlobalParam(
	name string,
	global int,
) {
	e.params[name] = &Param{
		global: global,
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

func (e *Emitter) SetFuncMinArgs(n int) {
	e.curFn().minArgs = n
}

func (e *Emitter) NewLabel() *Label {
	return &Label{
		fn:   e.curFn(),
		addr: -1,
	}
}

func (e *Emitter) ResolveLabel(label *Label) {
	if label.addr != -1 {
		panic("label already attached")
	}
	// The label is placed *after* the last instruction.
	instrs := e.curFn().instrs
	label.addr = len(instrs)
	for _, ref := range label.refs {
		instr := &instrs[ref]
		instr.operand1 = uint32(label.addr)
	}
}

func (e *Emitter) curFn() *Fn {
	return e.fnStack[len(e.fnStack)-1]
}

func (e *Emitter) Emit(pos token.Pos, op OpCode, operand1 uint32, operand2 uint16) {
	curFn := e.curFn()
	curFn.instrs = append(curFn.instrs, Instr{op, operand1, operand2})

	if pos.Filename != e.curFile {
		e.curFile = pos.Filename
		e.curFileID = e.AddString(pos.Filename)
	}

	var lastLoc *Location
	if len(curFn.locations) > 0 {
		lastLoc = &curFn.locations[len(curFn.locations)-1]
	}

	if lastLoc == nil ||
		lastLoc.filename != e.curFileID ||
		lastLoc.lineNum != pos.Line {
		curFn.locations = append(curFn.locations, Location{
			ip:       len(curFn.instrs) - 1,
			filename: e.curFileID,
			lineNum:  pos.Line,
		})
	}
}

func (e *Emitter) EmitJump(pos token.Pos, op OpCode, label *Label, operand2 uint16) {
	if label.addr == -1 {
		if e.curFn() != label.fn {
			panic("cannot jump across fns")
		}
		e.Emit(pos, op, 0, operand2)
		instrs := e.curFn().instrs
		lastEmittedInstr := len(instrs) - 1
		label.refs = append(label.refs, lastEmittedInstr)
	} else {
		e.Emit(pos, op, uint32(label.addr), 0)
	}
}

func (e *Emitter) AddExternalFunc(fn ExternFn) int {
	e.extFns = append(e.extFns, fn)
	return len(e.extFns) - 1
}

func (e *Emitter) AddString(s string) int {
	n, ok := e.stringMap[s]
	if !ok {
		n = e.AddLiteral(NewString(s))
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
		globals:   e.globals,
		fns:       e.fns,
		extFns:    e.extFns,
		literals:  e.literals,
		params:    e.params,
		reqParamN: e.reqParamN,
	}
}
