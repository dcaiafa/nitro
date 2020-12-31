package ast

type FuncStmt struct {
	funcBase

	Name   string
	Params ASTs
	Stmts  ASTs
}

func (a *FuncStmt) RunPass(ctx *Context, pass Pass) {

}
