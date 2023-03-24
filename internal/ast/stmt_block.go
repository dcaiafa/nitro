package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
)

type StmtBlock struct {
	PosImpl

	Stmts ASTs

	scope scope.Scope
}

var _ Scope = (*StmtBlock)(nil)

func (b *StmtBlock) Scope() scope.Scope {
	return b.scope
}

func (b *StmtBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = scope.NewScope(scope.Block)
	}
	ctx.RunPassChild(b, b.Stmts, pass)
}
