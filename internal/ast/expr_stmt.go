package ast

type ExprStmt struct {
	astBase

	Expr Expr
}

func (s *ExprStmt) RunPass(ctx *Context, pass Pass) {
	ctx.Push(s)
	s.Expr.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		//ctx.Emitter().Emit(runtime.OpEmit, 0)
	}
}
