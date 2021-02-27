package ast

import "github.com/dcaiafa/nitro/internal/symbol"

type FuncParam struct {
	astBase

	Name string

	sym *symbol.ParamSymbol
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
