package ast

type LambdaExpr struct {
	Func
}

func (e *LambdaExpr) isExpr() {}

func (e *LambdaExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		e.Func.DebugName = "$anonymous"
		e.Func.IsClosure = true
	}

	ctx.RunPassChild(e, &e.Func, pass)
}
