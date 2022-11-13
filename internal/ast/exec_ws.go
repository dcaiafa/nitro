package ast

type ExecWS struct {
	PosImpl
}

func (e *ExecWS) isExpr() {}

func (e *ExecWS) RunPass(ctx *Context, pass Pass) {
	// This AST is used as a placeholder in a transformation and will never be
	// evaluated.
	panic("not reached")
}
