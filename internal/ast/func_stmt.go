package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type FuncStmt struct {
	Func

	Name string

	sym types.Symbol
}

func (a *FuncStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		parentFn := ctx.CurrentFunc()

		if parentFn != nil {
			a.sym = parentFn.NewLocal()
		} else {
			a.sym = &types.FuncSymbol{}
		}

		a.sym.SetName(a.Name)
		a.sym.SetPos(a.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, a.sym) {
			return
		}
	}

	ctx.Push(a)
	a.Func.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case CreateAndResolveNames:
		if fnSym, ok := a.sym.(*types.FuncSymbol); ok {
			fnSym.FnNdx = a.Func.FnNdx
		}

	case Emit:
		if localSym, ok := a.sym.(*types.LocalVarSymbol); ok {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpPushLocalRef, uint16(localSym.LocalNdx), 0)
			emitter.Emit(runtime.OpMakeClosure, uint16(a.Func.FnNdx), 0)
			emitter.Emit(runtime.OpStore, 0, 0)
		}
	}
}
