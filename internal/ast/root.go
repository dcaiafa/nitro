package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Root struct {
	PosImpl

	Package Package

	externalFns ASTs

	rootScope *symbol.Scope
	globals   int
	metadata  *meta.Metadata
	initSym   *symbol.FuncSymbol
}

func (m *Root) AddNativeFn(name string, extFn *vm.NativeFn) {
	m.externalFns = append(
		m.externalFns,
		&ExternFn{
			Name:     name,
			ExternFn: extFn,
		})
}

func (m *Root) AddGlobalParam(ctx *Context, name string, param *meta.Param, pos token.Pos) symbol.Symbol {
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

func (m *Root) Scope() *symbol.Scope {
	return m.rootScope
}

func (m *Root) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateGlobals:
		m.rootScope = symbol.NewScope()
		m.metadata = new(meta.Metadata)
		m.initSym = &symbol.FuncSymbol{}
		m.initSym.SetReadOnly(true)
		m.initSym.SetName("$init")
		m.initSym.SetPos(m.Pos())
		if !m.rootScope.PutSymbol(ctx, m.initSym) {
			return
		}

	case Emit:
		emitter := ctx.Emitter()
		emitter.SetGlobalCount(m.globals)
		m.initSym.IdxFunc = emitter.NewFn(m.initSym.Name())
		emitter.PushFn(m.initSym.IdxFunc)
		emitter.SetFuncMinArgs(0)
		emitter.Emit(m.Pos(), vm.OpInitCallFrame, 0, 0)
		emitter.PopFn()
	}

	ctx.Push(m)
	m.externalFns.RunPass(ctx, pass)
	m.Package.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(m.initSym.IdxFunc)
		emitter.Emit(m.Pos(), vm.OpRet, 0, 0)
		emitter.PopFn()
	}
}

func (m *Root) Metadata() *meta.Metadata {
	return m.metadata
}
