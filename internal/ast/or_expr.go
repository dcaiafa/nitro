package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type OrExpr struct {
	astBase
	Left  Expr
	Right Expr
}

func (e *OrExpr) isExpr() {}

func (e *OrExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
	}

	ctx.RunPassChild(e, e.Left, pass)

	var skipLabel *runtime.Label
	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		skipLabel = emitter.NewLabel()
		emitter.Emit(runtime.OpDup, 0, 0)
		emitter.EmitJump(runtime.OpJumpIfTrue, skipLabel)
		emitter.Emit(runtime.OpPop, 1, 0)
	}

	ctx.RunPassChild(e, e.Right, pass)

	switch pass {
	case Emit:
		ctx.Emitter().ResolveLabel(skipLabel)
	}
}
