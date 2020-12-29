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
