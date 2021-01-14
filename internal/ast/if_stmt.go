package ast

type IfStmt struct {
	astBase
	Blocks ASTs
}

func (s *IfStmt) RunPass(ctx *Context, pass Pass) {
	for _, block := range s.Blocks {
		block.RunPass(ctx, pass)
	}
}

type IfBlock struct {
	astBase
	Pred  Expr
	Stmts ASTs
}

func (s *IfBlock) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}
