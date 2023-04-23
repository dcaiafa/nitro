package ast

import (
	"github.com/dcaiafa/nitro/internal/stub/analysis"
	"github.com/dcaiafa/nitro/internal/token"
)

type StructDecl struct {
	Name   string
	Fields ASTs
}

func (d *StructDecl) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(d, d.Fields, pass)

	if pass == Check {
		structDecl := &analysis.Struct{
			Name: d.Name,
		}
		for _, fieldAST := range d.Fields {
			fieldAST := fieldAST.(*StructField)
			field := &analysis.StructField{
				Name: fieldAST.Name,
				Type: fieldAST.Type.Type,
			}
			structDecl.Fields = append(structDecl.Fields, field)
		}

		err := ctx.Analysis.AddStruct(structDecl)
		if err != nil {
			ctx.Failf(token.Pos{}, "%v", err)
			return
		}
	}
}
