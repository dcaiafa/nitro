package ast

type AssignStmt struct {
	astBase
	Lvalue Expr
	Rvalue Expr
}

func (s *AssignStmt) RunPass(ctx *Context, pass Pass) {
	ctx.Push(s)
	s.Lvalue.RunPass(ctx, pass)
	s.Rvalue.RunPass(ctx, pass)
	ctx.Pop()

	if pass == CreateAndResolveNames {
		if _, ok := s.Lvalue.(LvalueExpr); !ok {
			ctx.Failf(
				s.Pos(),
				"Left side expression of assignment is not assignable. "+
					"The only assignable expression types are simple reference "+
					"(e.g. a = 1), and member access (e.g. a.b = 1)")
			return
		}
	}
}
