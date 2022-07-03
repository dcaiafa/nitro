package ast

import "github.com/dcaiafa/nitro/internal/symbol"

type StmtBlock struct {
	astBase

	Stmts ASTs

	scope *symbol.Scope
}

func (b *StmtBlock) Scope() *symbol.Scope {
	return b.scope
}

func (b *StmtBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = symbol.NewScope()
	}
	ctx.RunPassChild(b, b.Stmts, pass)
}
