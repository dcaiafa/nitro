package ast

import "github.com/dcaiafa/nitro/internal/vm"

type SliceExpr struct {
	astBase
	Target Expr
	Begin  Expr
	End    Expr
}

func (e *SliceExpr) isExpr() {}

func (e *SliceExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Target, pass)
	ctx.RunPassChild(e, e.Begin, pass)
	ctx.RunPassChild(e, e.End, pass)

	if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), vm.OpSlice, 0, 0)
	}
}
