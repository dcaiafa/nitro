package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
)

type LiteralExpr struct {
	astBase
	Val token.Token
}

func (a *LiteralExpr) isExpr() {}

func (a *LiteralExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		switch a.Val.Type {
		case token.Number:
			ctx.Emitter().Emit(runtime.OpPushInt, uint16(a.Val.Num), 0)

		default:
			panic("not implemented")
		}
	}
}
