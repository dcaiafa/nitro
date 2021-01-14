package ast

import "github.com/dcaiafa/nitro/internal/types"

type FuncParam struct {
	astBase

	Name string

	sym *types.ParamSymbol
}

func (a *FuncParam) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		parentFn := ctx.CurrentFunc()

		a.sym = parentFn.NewParam()
		a.sym.SetName(a.Name)
		a.sym.SetPos(a.Pos())

		if !parentFn.Scope().PutSymbol(ctx, a.sym) {
			return
		}
	}
}
