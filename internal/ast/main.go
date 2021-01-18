package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Main struct {
	astBase

	externalFns ASTs
	modules     ASTs

	rootScope *types.Scope
	globals   int
}

func (m *Main) AddExternalFn(name string, extFn runtime.ExternFn) {
	m.externalFns = append(
		m.externalFns,
		&ExternFn{
			Name:     name,
			ExternFn: extFn,
		})
}

func (m *Main) AddModule(module *Module) {
	m.modules = append(m.modules, module)
}

func (m *Main) NewGlobal() *types.GlobalVarSymbol {
	g := &types.GlobalVarSymbol{}
	g.GlobalNdx = m.globals
	m.globals++
	return g
}

func (m *Main) Scope() *types.Scope {
	return m.rootScope
}

func (m *Main) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		m.rootScope = types.NewScope()

	case Emit:
		ctx.Emitter().SetGlobalCount(m.globals)
	}

	ctx.Push(m)
	m.externalFns.RunPass(ctx, pass)
	m.modules.RunPass(ctx, pass)
	ctx.Pop()
}
