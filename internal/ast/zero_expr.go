package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ZeroExpr struct {
	astBase
}

func (e *ZeroExpr) isExpr() {}

func (e *ZeroExpr) RunPass(ctx *Context, pass Pass) {
	if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), runtime.OpNewInt, 0, 0)
	}
}
