package vm

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
) bool {
	if e.params[name] != nil {
		return false
	}
	e.params[name] = &Param{
		global: global,
	}
	return true
}

func (e *Emitter) NewFn(name string) int {
	idxName := e.AddLiteral(NewString(name))
	return e.AddLiteral(&Fn{name: idxName})
}

func (e *Emitter) PushFn(fn int) {
	e.fnStack = append(e.fnStack, e.literals[fn].(*Fn))
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
		instr.op1 = uint32(label.addr)
	}
}

func (e *Emitter) curFn() *Fn {
	return e.fnStack[len(e.fnStack)-1]
}

func (e *Emitter) Emit(pos token.Pos, op OpCode, operand1 uint32, operand2 uint16) {
	curFn := e.curFn()
	curFn.instrs = append(curFn.instrs, Instr{opc: op, op1: operand1, op2: operand2})

	if pos.Filename != e.curFile {
		e.curFile = pos.Filename
		e.curFileID = e.AddLiteral(NewString(pos.Filename))
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
			fn:       curFn.name,
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

func (e *Emitter) AddLiteral(v Value) int {
	if s, ok := v.(String); ok {
		n, ok := e.stringMap[s.String()]
		if !ok {
			n = e.addLiteral(s)
			e.stringMap[s.String()] = n
		}
		return n
	}
	return e.addLiteral(v)
}

func (e *Emitter) addLiteral(v Value) int {
	e.literals = append(e.literals, v)
	return len(e.literals) - 1
}

func (e *Emitter) SetGlobalCount(n int) {
	e.globals = n
}

func (e *Emitter) ToCompiledPackage() *CompiledPackage {
	return &CompiledPackage{
		globals:   e.globals,
		literals:  e.literals,
		params:    e.params,
		reqParamN: e.reqParamN,
	}
}
