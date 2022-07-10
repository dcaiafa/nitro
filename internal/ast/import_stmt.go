package ast

type Import struct {
	PosImpl

	Alias   string
	Package string
}

func (i *Import) RunPass(ctx *Context, pass Pass) {
}
