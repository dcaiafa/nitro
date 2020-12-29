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
		ctx.Emitter().Emit(runtime.OpCall, uint64(len(a.Args)))
	}
}
