package ast

import (
	"math"

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
		emitter := ctx.Emitter()

		switch a.Val.Type {
		case token.Int:
			if 0 <= a.Val.Int && a.Val.Int <= math.MaxUint16 {
				emitter.Emit(runtime.OpPushInt, uint16(a.Val.Int), 0)
			} else {
				literal := emitter.AddLiteral(runtime.Int(a.Val.Int))
				emitter.Emit(runtime.OpPushLiteral, uint16(literal), 0)
			}

		case token.String:
			str := emitter.AddString(a.Val.Str)
			emitter.Emit(runtime.OpPushLiteral, uint16(str), 0)

		default:
			panic("not implemented")
		}
	}
}
