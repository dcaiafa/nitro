package ast

type FuncStmt struct {
	Func
}

func (a *FuncStmt) RunPass(ctx *Context, pass Pass) {
}
