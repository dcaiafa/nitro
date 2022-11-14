package ast

import "github.com/dcaiafa/nitro/internal/vm"

type ArrayElement struct {
	PosImpl
	Val    Expr
	Expand bool
}

func (e *ArrayElement) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Val, pass)

	switch pass {
	case Emit:
		if e.Expand {
			ctx.Emitter().Emit(e.Pos(), vm.OpArrayExpandElemNoPop, 0, 0)
		} else {
			ctx.Emitter().Emit(e.Pos(), vm.OpArrayAppendNoPop, 0, 0)
		}
	}
}
