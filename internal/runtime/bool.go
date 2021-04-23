package runtime

import "strconv"

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
