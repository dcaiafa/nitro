package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

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
		a.sym.SetName(a.Name)
		a.sym.SetPos(a.Pos())
		a.sym.Scope = types.NewScope()
		a.sym.Fn = ctx.Emitter().NewFn()
		if !ctx.CurrentScope().PutSymbol(ctx, a.sym) {
			return
		}

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(a.sym.Fn)
		emitter.Emit(runtime.OpInitCallFrame, uint16(len(a.sym.Locals)), 0)
	}

	ctx.Push(a)
	a.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case CreateAndResolveNames:
		for i, local := range a.sym.Locals {
			local.Local = i
		}

	case Emit:
		synthRet := len(a.Stmts) == 0
		if !synthRet {
			_, hasReturnStmt := a.Stmts[len(a.Stmts)-1].(*ReturnStmt)
			synthRet = !hasReturnStmt
		}
		if synthRet {
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
