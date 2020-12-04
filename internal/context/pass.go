package context

type Pass int

const (
	Print Pass = iota

	CreateAndResolveNames
)

var Passes = []Pass{
	CreateAndResolveNames,
}

type PassRunner interface {
	RunPass(ctx *Context, pass Pass)
}
