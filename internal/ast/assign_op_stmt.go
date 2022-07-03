package ast

type AssignOperator int

type AssignOpStmt struct {
	PosImpl
	Op     Operator
	LValue *LValue
	RValue Expr

	rewritten AST
}

func (s *AssignOpStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Rewrite {
		// a += b    =>    a = a + b
		s.rewritten = &AssignStmt{
			Lvalues: ASTs{s.LValue},
			Rvalues: Exprs{&BinaryExpr{
				Left:  s.LValue.Expr,
				Op:    s.Op,
				Right: s.RValue,
			}},
		}
		s.rewritten.SetPos(s.Pos())
	}

	ctx.RunPassChild(s, s.rewritten, pass)
}
