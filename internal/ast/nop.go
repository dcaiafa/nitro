package ast

type Nop struct {
	PosImpl
}

func (n *Nop) RunPass(ctx *Context, pass Pass) {}
