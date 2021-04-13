package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
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
		f.sym.SetPos(token.Pos{
			Filename: "external-funcs",
		})
		f.sym.IdxFunc = ctx.Emitter().AddExternalFunc(f.ExternFn)
		if !ctx.CurrentScope().PutSymbol(ctx, f.sym) {
			return
		}
	}
}
