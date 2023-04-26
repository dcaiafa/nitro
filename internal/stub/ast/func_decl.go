package ast

import (
	"github.com/dcaiafa/nitro/internal/stub/analysis"
	"github.com/dcaiafa/nitro/internal/token"
)

type FuncDecl struct {
	Doc    *Doc
	ID     token.Token
	Params ASTs
	Rets   ASTs
}

func (d *FuncDecl) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(d, d.Params, pass)
	ctx.RunPassChild(d, d.Rets, pass)

	if pass == Check {
		sig := new(analysis.Signature)
		for i, paramAST := range d.Params {
			paramAST := paramAST.(*FuncParam)
			param := analysis.Param{
				Name: paramAST.ID.Str,
				Type: paramAST.Type.Type,
			}
			if paramAST.DefaultValue != nil {
				param.HasDefault = true
				param.Default = paramAST.DefaultValue.Value
			} else if len(sig.Params) > 0 && sig.Params[len(sig.Params)-1].HasDefault {
				ctx.Failf(token.Pos{}, "Only the last parameters are allowed to have default values")
				return
			}
			if paramAST.VarArg {
				if i != len(d.Params)-1 {
					ctx.Failf(token.Pos{}, "Only the last parameter can be vararg")
					return
				}
				param.VarArg = true
			}
			sig.Params = append(sig.Params, param)
		}
		for _, retAST := range d.Rets {
			retAST := retAST.(*TypeRef)
			sig.Rets = append(sig.Rets, retAST.Type)
		}
		fun := &analysis.Func{
			Name:       d.ID.Str,
			Signatures: []*analysis.Signature{sig},
		}
		err := ctx.Analysis.AddFunc(fun)
		if err != nil {
			ctx.Failf(token.Pos{}, "%v", err)
			return
		}
	}
}
