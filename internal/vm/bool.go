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
func (b Bool) Traits() Traits { return TraitEq }

func (b Bool) EvalOp(op Op, operand Value) (Value, error) {
	switch op {
	case OpEq:
		return NewBool(b == operand), nil
	default:
		return nil, fmt.Errorf("bool does not support this operation")
	}
}
