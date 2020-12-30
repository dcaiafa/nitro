package ast

type LValue struct {
	astBase

	Expr Expr
}

func (a *LValue) isExpr() {}

func (a *LValue) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		switch a.Expr.(type) {
		case *SimpleRef:
		case *MemberAccess:
		case *IndexExpr:
		default:
			ctx.Failf(a.Expr.Pos(), "Expression is not lvalue")
			return
		}
	}

	ctx.Push(a)
	a.Expr.RunPass(ctx, pass)
	ctx.Pop()
}
