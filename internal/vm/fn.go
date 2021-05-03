package vm

import "fmt"

type Fn struct {
	name      int
	locations []Location
	instrs    []Instr
	minArgs   int
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }

func (f *Fn) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}

func (f *Fn) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Function calls are handled directly by the VM.
	panic("not called")
}
