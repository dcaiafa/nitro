package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Root struct {
	PosImpl

	FuncRegistries []vm.ExportRegistry
	Package        Package

	rootScope symbol.Scope
	globals   int
	metadata  *meta.Metadata
}

var _ Scope = (*Root)(nil)

func (m *Root) AddGlobalParam(ctx *Context, name string, param *meta.Param, pos token.Pos) *symbol.GlobalVarSymbol {
	g := m.NewGlobal()
	g.SetName(name)
	g.SetPos(pos)

	if !ctx.CurrentScope().PutSymbol(ctx, g) {
		return nil
	}
	if !ctx.Emitter().AddGlobalParam(param.Name, g.GlobalNdx) {
		ctx.Failf(pos, "there is already a parameter named %q", name)
		return nil
	}
	m.metadata.Params = append(m.metadata.Params, param)
	return g
}

func (m *Root) AddUnit(module AST) {
	m.Package.Units = append(m.Package.Units, module)
}

func (m *Root) NewGlobal() *symbol.GlobalVarSymbol {
	g := &symbol.GlobalVarSymbol{}
	g.GlobalNdx = m.globals
	m.globals++
	return g
}

func (m *Root) Scope() symbol.Scope {
	return m.rootScope
}

func (m *Root) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateGlobals:
		m.rootScope = NewRootScope(ctx.Emitter(), "", m.FuncRegistries)
		m.metadata = new(meta.Metadata)

	case Emit:
		emitter := ctx.Emitter()
		emitter.SetGlobalCount(m.globals)
	}

	ctx.RunPassChild(m, &m.Package, pass)
}

func (m *Root) Metadata() *meta.Metadata {
	return m.metadata
}
