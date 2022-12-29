package ast

type Import struct {
	PosImpl

	Alias           string
	Package         string
	ExpandedPackage string
}

func (i *Import) RunPass(ctx *Context, pass Pass) {
}
