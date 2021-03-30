package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type FuncCallExpr struct {
	astBase
	Target Expr
	Args   Exprs
	RetN   int
}

func (c *FuncCallExpr) isExpr() {}

func (c *FuncCallExpr) RunPass(ctx *Context, pass Pass) {
	ctx.Push(c)
	c.Target.RunPass(ctx, pass)
	c.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		ctx.Emitter().Emit(c.Pos(), runtime.OpCall, uint32(len(c.Args)), uint16(c.RetN))
	}
}
