package vm

import (
	"github.com/dcaiafa/nitro/internal/meta"
)

type Program struct {
	Metadata  *meta.Metadata
	MainFnNdx int

	globals   int
	fns       []Fn
	extFns    []*NativeFn
	literals  []Value
	params    map[string]*Param
	reqParamN int
}
