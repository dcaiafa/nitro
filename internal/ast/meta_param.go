package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type MetaParam struct {
	astBase

	Name    string
	IsFlag  bool
	Default Expr
	Attribs ASTs

	globalSym symbol.Symbol
	param     *meta.Param
}

func (p *MetaParam) RunPass(ctx *Context, pass Pass) {
	var skip *vm.Label

	if pass == Check {
		p.param = &meta.Param{
			Name:   p.Name,
			IsFlag: p.IsFlag,
		}
	}

	if p.Default != nil {
		if pass == Emit {
			skip = ctx.Emitter().NewLabel()
			emitSymbolPush(p.Pos(), ctx.Emitter(), p.globalSym)
			ctx.Emitter().EmitJump(p.Pos(), vm.OpJumpIfTrue, skip, 0)
			emitSymbolRefPush(p.Pos(), ctx.Emitter(), p.globalSym)
		}

		ctx.RunPassChild(p, p.Default, pass)

		if pass == Emit {
			ctx.Emitter().Emit(p.Pos(), vm.OpStore, 1, 0)
			ctx.Emitter().ResolveLabel(skip)
		}
	}

	ctx.RunPassChild(p, p.Attribs, pass)

	if pass == Check {
		p.globalSym = ctx.Main().AddGlobalParam(ctx, p.Name, p.param, p.Pos())
		if p.globalSym == nil {
			return
		}
	}
}

type MetaAttrib struct {
	astBase

	Name  string
	Value vm.Value
}

func (a *MetaAttrib) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		paramAST := ctx.Parent().(*MetaParam)
		switch a.Name {
		case "name":
			nameStr, ok := a.Value.(vm.String)
			if !ok {
				ctx.Failf(a.Pos(), "'name' attribute value must be string")
				return
			}
			paramAST.param.Name = nameStr.String()

		case "type":
			typeStr, ok := a.Value.(vm.String)
			if !ok {
				ctx.Failf(a.Pos(), "'type' attribute value must be string")
				return
			}
			paramAST.param.Type = typeStr.String()

		case "desc":
			descStr, ok := a.Value.(vm.String)
			if !ok {
				ctx.Failf(a.Pos(), "'desc' attribute value must be string")
				return
			}
			paramAST.param.Desc = descStr.String()

		case "required":
			if a.Value == nil {
				paramAST.param.Required = true
			} else if reqBool, ok := a.Value.(vm.Bool); ok {
				paramAST.param.Required = reqBool.Bool()
			} else {
				ctx.Failf(a.Pos(), "'required' attribute value must be bool")
			}

		default:
			ctx.Failf(a.Pos(), "invalid attribute %q", a.Name)
		}
	}
}
