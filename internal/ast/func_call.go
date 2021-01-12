package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type FuncCall struct {
	astBase
	Target Expr
	Args   Exprs
}

func (s *FuncCall) isExpr() {}

func (a *FuncCall) RunPass(ctx *Context, pass Pass) {
	ctx.Push(a)
	a.Target.RunPass(ctx, pass)
	a.Args.RunPass(ctx, pass)
	ctx.Pop()

	if pass == Emit {
		retN := 0
		switch ctx.Parent().(type) {
		case Expr:
			retN = 1
		case *VarDeclStmt, *AssignStmt:
			// TODO: support multi-value assignments.
			retN = 1
		}

		ctx.Emitter().Emit(runtime.OpCall, uint16(len(a.Args)), byte(retN))
	}
}
