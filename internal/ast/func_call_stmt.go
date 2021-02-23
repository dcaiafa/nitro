package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type FuncCallStmt struct {
	astBase
	Target Expr
	Args   Exprs
}

func (c *FuncCallStmt) RunPass(ctx *Context, pass Pass) {
	ctx.Push(c)
	c.Target.RunPass(ctx, pass)
	c.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		ctx.Emitter().Emit(runtime.OpCall, uint16(len(c.Args)), 0)
	}
}
