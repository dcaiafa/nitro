package ast

import "github.com/dcaiafa/nitro/internal/vm"

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

	var skipLabel *vm.Label
	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		skipLabel = emitter.NewLabel()
		emitter.Emit(e.Pos(), vm.OpDup, 0, 0)
		emitter.EmitJump(e.Pos(), vm.OpJumpIfFalse, skipLabel, 0)
		emitter.Emit(e.Pos(), vm.OpPop, 1, 0)
	}

	ctx.RunPassChild(e, e.Right, pass)

	switch pass {
	case Emit:
		ctx.Emitter().ResolveLabel(skipLabel)
	}
}
