package ast

type StructField struct {
	Name string
	Type *TypeRef
}

func (f *StructField) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(f, f.Type, pass)
}
