package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Main struct {
	astBase

	externalFns ASTs
	modules     ASTs

	rootScope *symbol.Scope
	globals   int
	metadata  *meta.Metadata
}

func (m *Main) AddNativeFn(name string, extFn vm.NativeFn) {
	m.externalFns = append(
		m.externalFns,
		&ExternFn{
			Name:     name,
			ExternFn: extFn,
		})
}

func (m *Main) AddGlobalParam(ctx *Context, name string, param *meta.Param, pos token.Pos) symbol.Symbol {
	g := m.NewGlobal()
	g.SetName(name)
	g.SetPos(pos)
	if !m.rootScope.PutSymbol(ctx, g) {
		return nil
	}
	if !ctx.Emitter().AddGlobalParam(param.Name, g.GlobalNdx) {
		ctx.Failf(pos, "there is already a parameter named %q", name)
		return nil
	}
	m.metadata.Params = append(m.metadata.Params, param)
	return g
}

func (m *Main) AddModule(module AST) {
	m.modules = append(m.modules, module)
}

func (m *Main) NewGlobal() *symbol.GlobalVarSymbol {
	g := &symbol.GlobalVarSymbol{}
	g.GlobalNdx = m.globals
	m.globals++
	return g
}

func (m *Main) Scope() *symbol.Scope {
	return m.rootScope
}

func (m *Main) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		m.rootScope = symbol.NewScope()
		m.metadata = new(meta.Metadata)

	case Emit:
		ctx.Emitter().SetGlobalCount(m.globals)
	}

	ctx.Push(m)
	m.externalFns.RunPass(ctx, pass)
	m.modules.RunPass(ctx, pass)
	ctx.Pop()
}

func (m *Main) Metadata() *meta.Metadata {
	return m.metadata
}
