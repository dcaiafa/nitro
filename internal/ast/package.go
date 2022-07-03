package ast

import "github.com/dcaiafa/nitro/internal/symbol"

type Package struct {
	astBase

	Units ASTs

	scope *symbol.Scope
}

var _ Scope = (*Package)(nil)

func (p *Package) Scope() *symbol.Scope {
	return p.scope
}

func (p *Package) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		p.scope = symbol.NewScope()
	}
	ctx.RunPassChild(p, p.Units, pass)
}
