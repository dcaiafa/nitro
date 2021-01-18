package ast

type LValue struct {
	astBase

	Expr Expr
}

func (v *LValue) isExpr() {}

func (v *LValue) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		switch v.Expr.(type) {
		case *SimpleRef:
		case *MemberAccess:
		case *IndexExpr:
		default:
			ctx.Failf(v.Expr.Pos(), "Expression is not lvalue")
			return
		}
	}

	ctx.RunPassChild(v, v.Expr, pass)
}
