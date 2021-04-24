package vm

import (
	"fmt"
	"regexp"
)

type Regex struct {
	*regexp.Regexp
}

func (r *Regex) String() string { return r.Regexp.String() }
func (r *Regex) Type() string   { return "Regex" }

func (r *Regex) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("regex does not support this operation")
}

func (r *Regex) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("regex does not support this operation")
}

func NewRegex(r *regexp.Regexp) *Regex {
	return &Regex{Regexp: r}
}
