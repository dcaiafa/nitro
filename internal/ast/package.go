package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type Package struct {
	PosImpl

	Units ASTs

	scope *symbol.Scope

	main AST
}

var _ Scope = (*Package)(nil)

func (p *Package) Scope() *symbol.Scope {
	return p.scope
}

func (p *Package) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		callInit := &FuncCallExpr{
			PosImpl: p.PosImpl,
			Target: &SimpleRef{
				PosImpl: p.PosImpl,
				ID: token.Token{
					Pos:  p.Pos(),
					Type: token.String,
					Str:  "$init",
				},
			},
		}

		callMain := &FuncCallExpr{
			PosImpl: p.PosImpl,
			Target: &SimpleRef{
				PosImpl: p.PosImpl,
				ID: token.Token{
					Pos:  p.Pos(),
					Type: token.String,
					Str:  "main",
				},
			},
		}

		p.main = &FuncStmt{
			Name: "$main",
			Func: Func{
				PosImpl: p.PosImpl,
				Block: &StmtBlock{
					Stmts: ASTs{callInit, callMain},
				},
			},
		}

	case Check:
		p.scope = symbol.NewScope()
	}

	ctx.RunPassChild(p, p.Units, pass)
	ctx.RunPassChild(p, p.main, pass)
}
