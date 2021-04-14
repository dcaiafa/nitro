package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type MetaParam struct {
	astBase

	Name    string
	Default Expr

	globalSym symbol.Symbol
}

func (p *MetaParam) RunPass(ctx *Context, pass Pass) {
	var skip *runtime.Label

	if pass == Check {
		param := &meta.Param{
			Name: p.Name,
		}

		p.globalSym = ctx.Main().AddGlobalParam(ctx, param, p.Pos())
		if p.globalSym == nil {
			return
		}
	}

	if p.Default != nil {
		if pass == Emit {
			skip = ctx.Emitter().NewLabel()
			emitSymbolPush(p.Pos(), ctx.Emitter(), p.globalSym)
			ctx.Emitter().EmitJump(p.Pos(), runtime.OpJumpIfTrue, skip, 0)
			emitSymbolRefPush(p.Pos(), ctx.Emitter(), p.globalSym)
		}

		ctx.RunPassChild(p, p.Default, pass)

		if pass == Emit {
			ctx.Emitter().Emit(p.Pos(), runtime.OpStore, 1, 0)
			ctx.Emitter().ResolveLabel(skip)
		}
	}

}
