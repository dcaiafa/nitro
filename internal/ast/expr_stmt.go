package ast

type ExprStmt struct {
	astBase
	Expr Expr
}

func (c *ExprStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		switch expr := c.Expr.(type) {
		case *FuncCallExpr:
			expr.RetN = 0
		default:
			ctx.Failf(c.Pos(), "Expression not allowed without a statement")
		}
	}

	ctx.Push(c)
	c.Expr.RunPass(ctx, pass)
	ctx.Pop()
}
