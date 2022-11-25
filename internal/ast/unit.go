package ast

type Unit struct {
	PosImpl

	Meta    ASTs
	Imports ASTs
	Block   ASTs
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	// Validate that unit statements contain only variable and function
	// declarations. Other statement types are allowed by the parser to support
	// the simple script mode.
	if pass == Rewrite {
		for _, s := range u.Block {
			switch s.(type) {
			case *VarDeclStmt, *FuncStmt:
			default:
				ctx.Failf(s.Pos(), "Statement outside of a function is not allowed")
			}
		}
	}

	ctx.RunPassChild(u, u.Meta, pass)
	ctx.RunPassChild(u, u.Imports, pass)
	ctx.RunPassChild(u, u.Block, pass)
}

// SimpleScriptToPackage converts a simple script to the package format, if
// applicable.
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
