package ast

import (
	"math"

	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type LiteralExpr struct {
	PosImpl
	Val token.Token
}

func (e *LiteralExpr) isExpr() {}

func (e *LiteralExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		emitter := ctx.Emitter()

		switch e.Val.Type {
		case token.Nil:
			emitter.Emit(e.Pos(), vm.OpNil, 0, 0)

		case token.Int:
			if 0 <= e.Val.Int && e.Val.Int <= math.MaxUint16 {
				emitter.Emit(e.Pos(), vm.OpNewInt, uint32(e.Val.Int), 0)
			} else {
				literal := emitter.AddLiteral(vm.NewInt(e.Val.Int))
				emitter.Emit(e.Pos(), vm.OpLoadLiteral, uint32(literal), 0)
			}

		case token.Float:
			literal := emitter.AddLiteral(vm.NewFloat(e.Val.Float))
			emitter.Emit(e.Pos(), vm.OpLoadLiteral, uint32(literal), 0)

		case token.String:
			str := emitter.AddLiteral(vm.NewString(e.Val.Str))
			emitter.Emit(e.Pos(), vm.OpLoadLiteral, uint32(str), 0)

		case token.Bool:
			var v uint16 = 0
			if e.Val.Bool {
				v = 1
			}
			emitter.Emit(e.Pos(), vm.OpNewBool, uint32(v), 0)

		default:
			panic("not implemented")
		}
	}
}
