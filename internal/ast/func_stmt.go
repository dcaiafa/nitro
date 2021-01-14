package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type FuncStmt struct {
	Func

	Name string

	Sym types.Symbol
}

func (a *FuncStmt) RunPass(ctx *Context, pass Pass) {
	var funcSym *types.FuncSymbol

	switch pass {
	case CreateAndResolveNames:
		parentFn := ctx.CurrentFunc()

		if parentFn != nil {
			a.Sym = parentFn.NewLocal()
		} else {
			funcSym = &types.FuncSymbol{}
			funcSym.SetName(a.Name)
			funcSym.SetPos(a.Pos())
			if !ctx.CurrentScope().PutSymbol(ctx, funcSym) {
				return
			}
		}
	}

	ctx.Push(a)
	a.Func.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case CreateAndResolveNames:
		if funcSym != nil {
			funcSym.Fn = a.Func.FnNdx
		}

	case Emit:
		if localSym, ok := a.Sym.(*types.LocalVarSymbol); ok {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpPushLocalRef, uint16(localSym.Local), 0)
			emitter.Emit(runtime.OpMakeClosure, uint16(a.Func.FnNdx), 0)
			emitter.Emit(runtime.OpStore, 0, 0)
		}
	}
}
