package ast

type Nop struct {
	astBase
}

func (n *Nop) RunPass(ctx *Context, pass Pass) {}
