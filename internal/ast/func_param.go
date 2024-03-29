package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type FuncParam struct {
	PosImpl

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

	case Emit:
		if p.sym.Lifted() {
			ctx.Emitter().Emit(p.Pos(), vm.OpLiftArg, uint32(p.sym.ParamNdx), 0)
		}
	}
}
