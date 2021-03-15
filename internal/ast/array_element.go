package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ArrayElement struct {
	astBase
	Val Expr
}

func (e *ArrayElement) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Val, pass)

	switch pass {
	case Emit:
		ctx.Emitter().Emit(e.Pos(), runtime.OpArrayAppendNoPop, 0, 0)
	}
}
