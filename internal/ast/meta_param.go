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
	Attribs ASTs

	globalSym symbol.Symbol
	param     *meta.Param
}

func (p *MetaParam) RunPass(ctx *Context, pass Pass) {
	var skip *runtime.Label

	if pass == Check {
		p.param = &meta.Param{
			Name: p.Name,
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

	ctx.RunPassChild(p, p.Attribs, pass)

	if pass == Check {
		p.globalSym = ctx.Main().AddGlobalParam(ctx, p.param, p.Pos())
		if p.globalSym == nil {
			return
		}
	}
}

type MetaAttrib struct {
	astBase

	Name  string
	Value runtime.Value
}

func (a *MetaAttrib) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		paramAST := ctx.Parent().(*MetaParam)
		switch a.Name {
		case "type":
			typeStr, ok := a.Value.(runtime.String)
			if !ok {
				ctx.Failf(a.Pos(), "'type' attribute value must be string")
				return
			}
			paramAST.param.Type = typeStr.String()

		case "desc":
			descStr, ok := a.Value.(runtime.String)
			if !ok {
				ctx.Failf(a.Pos(), "'desc' attribute value must be string")
				return
			}
			paramAST.param.Desc = descStr.String()

		default:
			ctx.Failf(a.Pos(), "invalid attribute %q", a.Name)
		}
	}
}
