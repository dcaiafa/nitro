package vm

import (
	"fmt"
	"strconv"
)

var (
	True  = Bool{true}
	False = Bool{false}
)

type Bool struct {
	v bool
}

func NewBool(v bool) Bool {
	if v {
		return True
	} else {
		return False
	}
}

func (b Bool) Bool() bool     { return b.v }
func (b Bool) String() string { return strconv.FormatBool(b.v) }
func (b Bool) Type() string   { return "Bool" }

func (b Bool) EvalBinOp(op BinOp, operand Value) (Value, error) {
	if op == BinEq {
		return NewBool(b == operand), nil
	} else if op == BinNE {
		return NewBool(b != operand), nil
	} else {
		return nil, fmt.Errorf("bool does not support this operation")
	}
}

func (b Bool) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("array does not support this operation")
}
