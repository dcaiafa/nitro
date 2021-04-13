package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type FuncCallExpr struct {
	astBase
	Target Expr
	Args   Exprs
	RetN   int
	Expand bool
}

func (c *FuncCallExpr) isExpr() {}

func (c *FuncCallExpr) RunPass(ctx *Context, pass Pass) {
	ctx.Push(c)
	c.Target.RunPass(ctx, pass)
	c.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		operand1 := uint32(len(c.Args))
		if c.Expand {
			operand1 |= 0x80000000
		}
		ctx.Emitter().Emit(c.Pos(), runtime.OpCall, operand1, uint16(c.RetN))
	}
}
