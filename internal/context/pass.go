package context

type Pass int

const (
	Print Pass = iota

	ResolveNames
	CheckTypes
	Fold
	Emit
)

var Passes = []Pass{
	ResolveNames,
	CheckTypes,
	Fold,
	Emit,
}

type PassRunner interface {
	RunPass(ctx *Context, pass Pass) error
}
