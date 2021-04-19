package runtime

import "strconv"

type Bool struct {
	v bool
}

func NewBool(v bool) Bool { return Bool{v} }

func (b Bool) Bool() bool     { return b.v }
func (b Bool) String() string { return strconv.FormatBool(b.v) }
func (b Bool) Type() string   { return "Bool" }
