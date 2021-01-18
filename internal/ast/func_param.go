package ast

import "github.com/dcaiafa/nitro/internal/types"

type FuncParam struct {
	astBase

	Name string

	sym *types.ParamSymbol
}

func (p *FuncParam) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		parentFn := ctx.CurrentFunc()

		p.sym = parentFn.NewParam()
		p.sym.SetName(p.Name)
		p.sym.SetPos(p.Pos())

		if !parentFn.Scope().PutSymbol(ctx, p.sym) {
			return
		}
	}
}
