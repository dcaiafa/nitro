package ast

import (
	"github.com/dcaiafa/nitro/internal/context"
	"github.com/dcaiafa/nitro/internal/types"
)

type Module struct {
	astBase
	Stmts ASTs

	sym *types.Symbol
}

func (p *Module) RunPass(ctx *context.Context, pass context.Pass) {
	if p.sym == nil {
		p.sym = types.NewSymbol("")
		p.sym.Type = types.Module
		p.sym.Scope = types.NewScope()
	}

	ctx.Scopes().Push(p.sym.Scope)
	p.runPassChildren(ctx, pass)
	ctx.Scopes().Pop()
}

func (p *Module) runPassChildren(ctx *context.Context, pass context.Pass) {
	p.Stmts.RunPass(ctx, pass)
}
