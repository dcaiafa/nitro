package ast

import "github.com/dcaiafa/nitro/internal/vm"

type ArrayElement struct {
	PosImpl
	Val Expr
}

func (e *ArrayElement) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Val, pass)

	switch pass {
	case Emit:
		ctx.Emitter().Emit(e.Pos(), vm.OpArrayAppendNoPop, 0, 0)
	}
}
