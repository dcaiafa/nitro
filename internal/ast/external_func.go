package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type ExternFn struct {
	astBase

	Name     string
	ExternFn runtime.ExternFn

	sym *types.FuncSymbol
}

func (f *ExternFn) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		f.sym = &types.FuncSymbol{}
		f.sym.SetName(f.Name)
		f.sym.External = true
		f.sym.Fn = ctx.Emitter().AddExternalFunc(f.ExternFn)
		if !ctx.CurrentScope().PutSymbol(ctx, f.sym) {
			return
		}
	}
}
