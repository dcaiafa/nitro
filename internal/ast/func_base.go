package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Func struct {
	astBase

	Name   string
	Params ASTs
	Stmts  ASTs

	Sym *types.FuncSymbol
}

func (a *Func) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		a.Sym = &types.FuncSymbol{}
		a.Sym.SetName(a.Name)
		a.Sym.SetPos(a.Pos())
		a.Sym.Scope = types.NewScope()
		a.Sym.Fn = ctx.Emitter().NewFn()
		if !ctx.CurrentScope().PutSymbol(ctx, a.Sym) {
			return
		}

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(a.Sym.Fn)
		emitter.Emit(runtime.OpInitCallFrame, uint16(len(a.Sym.Locals)), 0)
	}

	ctx.Push(a)
	a.Params.RunPass(ctx, pass)
	a.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case CreateAndResolveNames:
		for i, local := range a.Sym.Locals {
			local.Local = i
		}
		for i, param := range a.Sym.Params {
			param.Arg = i
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

func (a *Func) Scope() *types.Scope {
	return a.Sym.Scope
}

func (a *Func) Func() *types.FuncSymbol {
	return a.Sym
}
