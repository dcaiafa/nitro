package ast

type Unit struct {
	PosImpl

	Meta    ASTs
	Imports ASTs
	Block   ASTs
}

func (m *Unit) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(m, m.Meta, pass)
	ctx.RunPassChild(m, m.Imports, pass)
	ctx.RunPassChild(m, m.Block, pass)
}

func SimpleScriptToPackage(u *Unit) {
	main := &FuncStmt{
		Name: "main",
		Func: Func{
			Block: &StmtBlock{
				Stmts: u.Block,
			},
		},
	}

	main.SetPos(u.Block.Pos())
	u.Block = ASTs{main}
}
