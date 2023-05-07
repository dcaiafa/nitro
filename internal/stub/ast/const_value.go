package ast

type ConstValue struct {
	Expr string
}

func (v *ConstValue) RunPass(ctx *Context, pass Pass) {}
