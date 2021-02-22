package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ArrayLiteral struct {
	astBase
	Elements ASTs
}

func (a *ArrayLiteral) isExpr() {}

func (a *ArrayLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		ctx.Emitter().Emit(runtime.OpNewArray, 0, 0)
	}

	ctx.RunPassChild(a, a.Elements, pass)
}
