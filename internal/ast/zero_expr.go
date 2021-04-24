package ast

import "github.com/dcaiafa/nitro/internal/vm"

type ZeroExpr struct {
	astBase
}

func (e *ZeroExpr) isExpr() {}

func (e *ZeroExpr) RunPass(ctx *Context, pass Pass) {
	if pass == Emit {
		ctx.Emitter().Emit(e.Pos(), vm.OpNewInt, 0, 0)
	}
}
