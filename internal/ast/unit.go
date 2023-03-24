package ast

import "github.com/dcaiafa/nitro/internal/scope"

type Unit struct {
	PosImpl

	scope *scope.SimpleScope

	Meta    ASTs
	Imports ASTs
	Block   ASTs
}

var _ Scope = (*Unit)(nil)

func (u *Unit) Scope() scope.Scope {
	return u.scope
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		// Validate that unit statements contain only variable and function
		// declarations. Other statement types are allowed by the parser to support
		// the simple script mode.
		for _, s := range u.Block {
			switch s.(type) {
			case *VarDeclStmt, *FuncStmt:
			default:
				ctx.Failf(s.Pos(), "Statement outside of a function is not allowed")
			}
		}

	case CreateGlobals:
		u.scope = scope.NewScope(scope.Unit)
	}

	ctx.RunPassChild(u, u.Meta, pass)
	ctx.RunPassChild(u, u.Imports, pass)
	ctx.RunPassChild(u, u.Block, pass)
}

func (u *Unit) ConvertSimple() {
	// If the script has a "main" function, then it is already assumed to be in
	// the strict format.
	for _, s := range u.Block {
		if funcStmt, ok := s.(*FuncStmt); ok {
			if funcStmt.Name == "main" {
				return
			}
		}
	}

	// Script is in simple format. Wrap all statements in a synthesized "main"
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
