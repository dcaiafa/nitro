package ast

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type FuncStmt struct {
	Func

	Name string

	Sym *types.LocalVarSymbol
}

func (a *FuncStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		a.Sym = &types.LocalVarSymbol{}
		a.Sym.SetName(a.Name)
		a.Sym.SetPos(a.Pos())

		fn := ctx.CurrentFunc().Sym
		fn.ClosureN++
		a.Func.Name = fmt.Sprintf("$func%d", fn.ClosureN)

		if !ctx.CurrentScope().PutSymbol(ctx, a.Sym) {
			return
		}

		fn.Locals = append(fn.Locals, a.Sym)
	}

	ctx.Push(a)
	a.Func.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.Emit(runtime.OpPushLocalRef, uint16(a.Sym.Local), 0)
		emitter.Emit(runtime.OpMakeClosure, uint16(a.Func.Sym.Fn), 0)
		emitter.Emit(runtime.OpStore, 0, 0)
	}
}
