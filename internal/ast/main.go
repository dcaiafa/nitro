package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type Main struct {
	astBase

	externalFns ASTs
	modules     ASTs

	rootScope *symbol.Scope
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

func (m *Main) AddGlobalParam(ctx *Context, param *meta.Param, pos token.Pos) bool {
	g := m.NewGlobal()
	g.SetName(param.Name)
	g.SetPos(pos)
	if !m.Scope().PutSymbol(ctx, g) {
		return false
	}

	ctx.Emitter().AddGlobalParam(
		param.Name,
		g.GlobalNdx,
		param.Type,
		param.Required,
		param.Default)

	return true
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

	case Emit:
		ctx.Emitter().SetGlobalCount(m.globals)
	}

	ctx.Push(m)
	m.externalFns.RunPass(ctx, pass)
	m.modules.RunPass(ctx, pass)
	ctx.Pop()
}
