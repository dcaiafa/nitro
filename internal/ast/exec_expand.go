package ast

type ExecExpand struct {
  PosImpl
  Expr Expr
}

func (e *ExecExpand) isExpr() {}

func (e *ExecExpand) RunPass(ctx *Context, pass Pass) {
  // This AST is used as a placeholder in a transformation and will never be
  // evaluated.
  panic("not reached")
}
