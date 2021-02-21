package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ObjectField struct {
	astBase

	NameID   string
	NameExpr Expr
	Val      Expr
}

func (s *ObjectField) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.NameExpr, pass)

	switch pass {
	case Emit:
		if s.NameID != "" {
			emitter := ctx.Emitter()
			nameID := emitter.AddString(s.NameID)
			emitter.Emit(runtime.OpLoadLiteral, uint16(nameID), 0)
		}
	}

	ctx.RunPassChild(s, s.Val, pass)

	switch pass {
	case Emit:
		ctx.Emitter().Emit(runtime.OpObjectPutNoPop, 0, 0)
	}
}
