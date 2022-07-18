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

func (r *Regex) EvalOp(op Op, operand Value) (Value, error) {
  if op == OpEq {
    return NewBool(r == operand), nil;
  }
	return nil, fmt.Errorf("regex does not support this operation")
}

func NewRegex(r *regexp.Regexp) *Regex {
	return &Regex{Regexp: r}
}
