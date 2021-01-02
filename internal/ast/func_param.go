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
		a.sym = &types.ParamSymbol{}
		a.sym.SetName(a.Name)
		a.sym.SetPos(a.Pos())

		ownerFn := ctx.CurrentFunc()
		ownerFn.Scope().PutSymbol(ctx, a.sym)
		ownerFn.Sym.Params = append(ownerFn.Sym.Params, a.sym)
	}
}
