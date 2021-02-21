package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ObjectLiteral struct {
	astBase
	Fields ASTs
}

func (s *ObjectLiteral) isExpr() {}

func (s *ObjectLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		ctx.Emitter().Emit(runtime.OpNewObject, 0, 0)
	}

	ctx.RunPassChild(s, s.Fields, pass)
}
