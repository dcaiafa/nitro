package ast

import (
	"github.com/dcaiafa/nitro/internal/meta"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type MetaParam struct {
	PosImpl

	Name    string
	IsFlag  bool
	Default Expr
	Attribs ASTs

	init *MetaParamInit

	globalSym *symbol.GlobalVarSymbol
	param     *meta.Param
}

func (p *MetaParam) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		if p.Default != nil {
			p.init = &MetaParamInit{
				PosImpl: p.PosImpl,
				Default: p.Default,
			}
			ctx.Package().AddInitStmt(p.init)
			p.Default = nil
		}

	case CreateGlobals:
		// Initialize `param` here so that MetaAttrib (in Attribs) can access, and
		// potentially change it.
		p.param = &meta.Param{
			Name:   p.Name,
			IsFlag: p.IsFlag,
		}
	}

	ctx.RunPassChild(p, p.Attribs, pass)

	switch pass {
	case CreateGlobals:
		p.globalSym = ctx.Main().AddGlobalParam(ctx, p.Name, p.param, p.Pos())
		if p.globalSym == nil {
			return
		}
		if p.init != nil {
			p.init.GlobalSym = p.globalSym
		}
	}
}

type MetaAttrib struct {
	PosImpl

	Name  string
	Value vm.Value
}

func (a *MetaAttrib) RunPass(ctx *Context, pass Pass) {
	if pass == CreateGlobals {
		// Attribute evaluation must happen at CreateGlobals because "name" can
		// override the name of the global parameter which must be created at
		// CreateGlobals.

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

type MetaParamInit struct {
	PosImpl
	Default   Expr
	GlobalSym *symbol.GlobalVarSymbol
}

func (p *MetaParamInit) RunPass(ctx *Context, pass Pass) {
	var skip *vm.Label

	if pass == Emit {
		emitter := ctx.Emitter()
		skip = emitter.NewLabel()
		emitSymbolPush(p.Pos(), emitter, p.GlobalSym)
		emitter.EmitJump(p.Pos(), vm.OpJumpIfTrue, skip, 0)
		emitSymbolRefPush(p.Pos(), emitter, p.GlobalSym)
	}

	ctx.RunPassChild(p, p.Default, pass)

	if pass == Emit {
		emitter := ctx.Emitter()
		emitter.Emit(p.Pos(), vm.OpStore, 1, 0)
		emitter.ResolveLabel(skip)
	}
}
