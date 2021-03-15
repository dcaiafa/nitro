package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type AndExpr struct {
	astBase
	Left  Expr
	Right Expr
}

func (e *AndExpr) isExpr() {}

func (e *AndExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
	}

	ctx.RunPassChild(e, e.Left, pass)

	var skipLabel *runtime.Label
	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		skipLabel = emitter.NewLabel()
		emitter.Emit(e.Pos(), runtime.OpDup, 0, 0)
		emitter.EmitJump(e.Pos(), runtime.OpJumpIfFalse, skipLabel)
		emitter.Emit(e.Pos(), runtime.OpPop, 1, 0)
	}

	ctx.RunPassChild(e, e.Right, pass)

	switch pass {
	case Emit:
		ctx.Emitter().ResolveLabel(skipLabel)
	}
}
