package ast

import "github.com/dcaiafa/nitro/internal/token"

type ImportDecl struct {
	Alias  string
	Import string
}

func (d *ImportDecl) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		err := ctx.Analysis.AddImport(d.Alias, d.Import)
		if err != nil {
			ctx.Failf(token.Pos{}, "%v", err)
			return
		}
	}
}
