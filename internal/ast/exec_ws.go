package ast

type ExecWS struct {
	PosImpl
}

func (e *ExecWS) isExpr() {}

func (e *ExecWS) RunPass(ctx *Context, pass Pass) {}
