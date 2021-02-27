package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type ExternFn struct {
	astBase

	Name     string
	ExternFn runtime.ExternFn

	sym *symbol.FuncSymbol
}

func (f *ExternFn) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		f.sym = &symbol.FuncSymbol{}
		f.sym.SetName(f.Name)
		f.sym.External = true
		f.sym.FnNdx = ctx.Emitter().AddExternalFunc(f.ExternFn)
		if !ctx.CurrentScope().PutSymbol(ctx, f.sym) {
			return
		}
	}
}
