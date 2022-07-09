package ast

type Import struct {
	PosImpl

	Alias      string
	ModuleName string
}

func (i *Import) RunPass(ctx *Context, pass Pass) {
}
