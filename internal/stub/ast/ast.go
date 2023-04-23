package ast

type AST interface {
	RunPass(ctx *Context, pass Pass)
}

type ASTs []AST

func (asts ASTs) RunPass(ctx *Context, pass Pass) {
	for _, ast := range asts {
		ast.RunPass(ctx, pass)
	}
}
