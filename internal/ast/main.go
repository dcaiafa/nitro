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

func (m *Main) AddExternalFn(name string, extFn runtime.ExternalFunc) {
	m.externalFns = append(m.externalFns, &ExternalFunc{
		Name:         name,
		ExternalFunc: extFn,
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

type funcBase struct {
	astBase
	Name  string
	Stmts ASTs

	sym *types.FuncSymbol
}

func (a *funcBase) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		a.sym = &types.FuncSymbol{}
		a.sym.Scope = types.NewScope()
		a.sym.Fn = ctx.Emitter().NewFn()

	case Emit:
		ctx.Emitter().PushFn(a.sym.Fn)
	}

	ctx.Push(a)
	a.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		synthesizeReturn := len(a.Stmts) == 0
		if !synthesizeReturn {
			_, hasReturnStmt := a.Stmts[len(a.Stmts)-1].(*ReturnStmt)
			synthesizeReturn = !hasReturnStmt
		}
		if synthesizeReturn {
			ctx.Emitter().Emit(runtime.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}

func (a *funcBase) Scope() *types.Scope {
	return a.sym.Scope
}

func (a *funcBase) Func() *types.FuncSymbol {
	return a.sym
}
