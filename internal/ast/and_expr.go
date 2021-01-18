package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type AndExpr struct {
	astBase
	Left  Expr
	Op    Operator
	Right Expr
}

func (a *AndExpr) isExpr() {}

func (a *AndExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
	}

	ctx.RunPassChild(a, a.Left, pass)

	var skipLabel *runtime.Label
	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		skipLabel = emitter.NewLabel()
		emitter.Emit(runtime.OpDup, 0, 0)
		emitter.EmitJump(runtime.OpJumpIfFalse, skipLabel)
		emitter.Emit(runtime.OpPop, 0, 0)
	}

	ctx.RunPassChild(a, a.Right, pass)

	switch pass {
	case Emit:
		ctx.Emitter().AssignLabel(skipLabel)
	}
}
