package ast

import (
	"github.com/dcaiafa/nitro/internal/stub/analysis"
	"github.com/dcaiafa/nitro/internal/token"
)

type TypeDecl struct {
	Doc    *Doc
	ID     token.Token
	GoType *GoType
}

func (d *TypeDecl) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		typ := analysis.Type{
			Name: d.ID.Str,
			GoType: analysis.GoType{
				Ref:     d.GoType.Ref,
				Package: d.GoType.Package.Str,
				Name:    d.GoType.ID.Str,
			},
		}
		err := ctx.Analysis.AddGoType(typ)
		if err != nil {
			ctx.Failf(token.Pos{}, "%s", err.Error())
			return
		}
	}
}
