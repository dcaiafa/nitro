package ast

import "github.com/dcaiafa/nitro/internal/vm"

type FuncCallExpr struct {
	PosImpl
	Target   Expr
	Args     Exprs
	RetN     int
	Expand   bool
	Pipeline bool
}

func (c *FuncCallExpr) isExpr() {}

func (c *FuncCallExpr) RunPass(ctx *Context, pass Pass) {
	ctx.Push(c)
	c.Target.RunPass(ctx, pass)
	c.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		if len(c.Args) > int(vm.CallArgCountMask) {
			ctx.Failf(c.Pos(), "Too many arguments")
			return
		}
		operand1 := uint32(len(c.Args))
		if c.Expand {
			operand1 |= vm.CallExpandFlag
		}
		if c.Pipeline {
			operand1 |= vm.CallPipelineFlag
		}
		ctx.Emitter().Emit(c.Pos(), vm.OpCall, operand1, uint16(c.RetN))
	}
}
