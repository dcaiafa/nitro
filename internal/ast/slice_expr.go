package ast

import "github.com/dcaiafa/nitro/internal/vm"

type SliceExpr struct {
	PosImpl
	Target Expr
	Begin  Expr
	End    Expr
}

func (e *SliceExpr) isExpr() {}

func (e *SliceExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Target, pass)

	if e.Begin != nil {
		ctx.RunPassChild(e, e.Begin, pass)
	} else if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), vm.OpNewInt, 0, 0)
	}
	if e.End != nil {
		ctx.RunPassChild(e, e.End, pass)
	} else if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), vm.OpNewInt, 0xFFFFFFFF, 0)
	}
	if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), vm.OpSlice, 0, 0)
	}
}
