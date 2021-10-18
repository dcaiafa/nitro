package ast

type AnonFuncExpr struct {
	Func
}

func (e *AnonFuncExpr) isExpr() {}

func (e *AnonFuncExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		e.DebugName = "$anon"
		e.IsClosure = true
	}
	ctx.RunPassChild(e, &e.Func, pass)
}
