package runtime

import "regexp"

type Regex struct {
	*regexp.Regexp
}

func (r *Regex) String() string { return r.Regexp.String() }
func (r *Regex) Type() string   { return "Regex" }

func NewRegex(r *regexp.Regexp) *Regex {
	return &Regex{Regexp: r}
}
