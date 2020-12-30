package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type AssignStmt struct {
	astBase
	Lvalue *LValue
	Rvalue Expr
}

func (s *AssignStmt) RunPass(ctx *Context, pass Pass) {
	ctx.Push(s)
	s.Lvalue.RunPass(ctx, pass)
	s.Rvalue.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.Emit(runtime.OpStore, 0, 0)
	}
}
