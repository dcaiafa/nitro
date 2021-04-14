package runtime

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type Program struct {
	Metadata *meta.Metadata

	globals   int
	fns       []Fn
	extFns    []ExternFn
	literals  []Value
	params    map[string]*Param
	reqParamN int
}
