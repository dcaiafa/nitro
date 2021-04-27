package vm

import "fmt"

type NativeFn func(m *VM, args []Value, nRet int) ([]Value, error)

func (f NativeFn) String() string { return "<func>" }
func (f NativeFn) Type() string   { return "Func" }
func (f NativeFn) isCallable()    {}

func (f NativeFn) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}

func (f NativeFn) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("func does not support this operation")
}
