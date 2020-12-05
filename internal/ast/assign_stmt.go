package ast

import "github.com/dcaiafa/nitro/internal/context"

type AssignStmt struct {
	astBase
	Lvalue Expr
	Rvalue Expr
}

func (s *AssignStmt) RunPass(ctx *context.Context, pass context.Pass) {
	s.runPassChildren(ctx, pass)

	if pass == context.CreateAndResolveNames {
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

func (s *AssignStmt) runPassChildren(ctx *context.Context, pass context.Pass) {
	s.Lvalue.RunPass(ctx, pass)
	s.Rvalue.RunPass(ctx, pass)
}
