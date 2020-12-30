package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Main struct {
	funcBase

	globalScope *types.Scope
	externalFns ASTs
}

func (m *Main) AddExternalFn(name string, extFn runtime.ExternFn) {
	m.externalFns = append(
		m.externalFns,
		&ExternFn{
			Name:     name,
			ExternFn: extFn,
		})
}

func (m *Main) Scope() *types.Scope {
	return m.globalScope
}

func (m *Main) RunPass(ctx *Context, pass Pass) {
	if pass == CreateAndResolveNames {
		m.Name = "$main"
		m.globalScope = types.NewScope()
	}

	ctx.Push(m)
	m.externalFns.RunPass(ctx, pass)
	m.funcBase.RunPass(ctx, pass)
	ctx.Pop()
}
