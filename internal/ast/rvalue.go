package ast

type RValue struct {
	astBase

	Expr Expr
}

func (a *RValue) RunPass(ctx *Context, pass Pass) {
	ctx.Push(a)
	a.Expr.RunPass(ctx, pass)
	ctx.Pop()
}
