package ast

type ConstValue struct {
	Value any
}

func (v *ConstValue) RunPass(ctx *Context, pass Pass) {
}
