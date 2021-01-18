package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type FuncCall struct {
	astBase
	Target Expr
	Args   Exprs
}

func (c *FuncCall) isExpr() {}

func (c *FuncCall) RunPass(ctx *Context, pass Pass) {
	ctx.Push(c)
	c.Target.RunPass(ctx, pass)
	c.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		retN := 1
		switch ctx.Parent().(type) {
		case *ExprStmt:
			retN = 0
		case *VarDeclStmt, *AssignStmt:
			// TODO: support multi-value assignments.
			retN = 1
		}

		ctx.Emitter().Emit(runtime.OpCall, uint16(len(c.Args)), byte(retN))
	}
}
