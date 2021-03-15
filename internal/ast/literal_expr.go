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

func (e *LiteralExpr) isExpr() {}

func (e *LiteralExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		emitter := ctx.Emitter()

		switch e.Val.Type {
		case token.Int:
			if 0 <= e.Val.Int && e.Val.Int <= math.MaxUint16 {
				emitter.Emit(e.Pos(), runtime.OpNewInt, uint16(e.Val.Int), 0)
			} else {
				literal := emitter.AddLiteral(runtime.NewInt(e.Val.Int))
				emitter.Emit(e.Pos(), runtime.OpLoadLiteral, uint16(literal), 0)
			}

		case token.Float:
			literal := emitter.AddLiteral(runtime.NewFloat(e.Val.Float))
			emitter.Emit(e.Pos(), runtime.OpLoadLiteral, uint16(literal), 0)

		case token.String:
			str := emitter.AddString(e.Val.Str)
			emitter.Emit(e.Pos(), runtime.OpLoadLiteral, uint16(str), 0)

		case token.Bool:
			var v uint16 = 0
			if e.Val.Bool {
				v = 1
			}
			emitter.Emit(e.Pos(), runtime.OpNewBool, uint16(v), 0)

		default:
			panic("not implemented")
		}
	}
}
