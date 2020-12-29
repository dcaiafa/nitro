package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type ExternalFunc struct {
	astBase

	Name         string
	ExternalFunc runtime.ExternalFunc

	sym *types.FuncSymbol
}

func (f *ExternalFunc) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		f.sym = &types.FuncSymbol{}
		f.sym.SetName(f.Name)
		f.sym.External = true
		f.sym.Fn = ctx.Emitter().AddExternalFunc(f.ExternalFunc)
		ctx.CurrentScope().PutSymbol(f.sym)
	}
}
