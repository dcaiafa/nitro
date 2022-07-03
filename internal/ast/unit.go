package ast

import "github.com/dcaiafa/nitro/internal/token"

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
	callInit := &FuncCallExpr{
		Target: &SimpleRef{
			ID: token.Token{
				Pos:  u.Pos(),
				Type: token.String,
				Str:  "$init",
			},
		},
	}

	main := &FuncStmt{
		Name: "$main",
		Func: Func{
			Block: &StmtBlock{
				Stmts: append(ASTs{callInit}, u.Block...),
			},
		},
	}

	main.SetPos(u.Block.Pos())
	u.Block = ASTs{main}
}
