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
	// If the script has a "main" function, then it is already in package format.
	for _, s := range u.Block {
		if funcStmt, ok := s.(*FuncStmt); ok {
			if funcStmt.Name == "main" {
				return
			}
		}
	}

	// Script is in simple form. Wrap all statements in a synthesized "main"
	// function.
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
