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
func (f *Fn) isCallable()    {}

func (f *Fn) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}

func (f *Fn) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}
