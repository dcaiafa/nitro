package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type TernaryExpr struct {
	astBase

	Condition Expr
	Then      Expr
	Else      Expr
}

func (e *TernaryExpr) isExpr() {}

func (e *TernaryExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Condition, pass)

	var elseLabel *runtime.Label
	var endLabel *runtime.Label
	if pass == Emit {
		emitter := ctx.Emitter()
		elseLabel = emitter.NewLabel()
		endLabel = emitter.NewLabel()
		emitter.EmitJump(e.Pos(), runtime.OpJumpIfFalse, elseLabel, 0)
	}

	ctx.RunPassChild(e, e.Then, pass)
	if pass == Emit {
		emitter := ctx.Emitter()
		emitter.EmitJump(e.Pos(), runtime.OpJump, endLabel, 0)
		emitter.ResolveLabel(elseLabel)
	}

	ctx.RunPassChild(e, e.Else, pass)

	if pass == Emit {
		ctx.Emitter().ResolveLabel(endLabel)
	}
}
