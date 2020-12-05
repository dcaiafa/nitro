package ast

import (
	"github.com/dcaiafa/nitro/internal/context"
	"github.com/dcaiafa/nitro/internal/typecheck"
)

type Module struct {
	astBase
	Stmts ASTs

	scope *typecheck.Scope
}

func (p *Module) RunPass(ctx *context.Context, pass context.Pass) {
	if p.scope == nil {
		p.scope = typecheck.NewScope()
	}

	ctx.Scopes().Push(p.scope)
	p.runPassChildren(ctx, pass)
	ctx.Scopes().Pop()
}

func (p *Module) runPassChildren(ctx *context.Context, pass context.Pass) {
	p.Stmts.RunPass(ctx, pass)
}
