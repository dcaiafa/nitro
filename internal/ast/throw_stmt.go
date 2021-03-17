package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type ThrowStmt struct {
	astBase

	Expr Expr
}

func (s *ThrowStmt) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.Expr, pass)

	if pass == Emit {
		ctx.Emitter().Emit(s.Pos(), runtime.OpThrow, 0, 0)
	}
}
