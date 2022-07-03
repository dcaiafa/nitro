package ast

import "github.com/dcaiafa/nitro/internal/token"

type IncDecOp int

const (
	IncDecOpUndefined IncDecOp = iota
	IncDecOpInc
	IncDecOpDec
)

type IncDecStmt struct {
	PosImpl
	LValue *LValue
	Op     IncDecOp

	rewritten AST
}

func (s *IncDecStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Rewrite {
		// a++   =>  a = a + 1
		var binOp Operator
		switch s.Op {
		case IncDecOpInc:
			binOp = BinOpPlus
		case IncDecOpDec:
			binOp = BinOpMinus
		default:
			panic("unreachable")
		}

		s.rewritten = &AssignStmt{
			Lvalues: ASTs{s.LValue},
			Rvalues: Exprs{
				&BinaryExpr{
					Left: s.LValue.Expr,
					Op:   binOp,
					Right: &LiteralExpr{
						Val: token.Token{
							Type: token.Int,
							Int:  1,
						},
					},
				},
			},
		}

		s.rewritten.SetPos(s.Pos())
	}

	ctx.RunPassChild(s, s.rewritten, pass)
}
