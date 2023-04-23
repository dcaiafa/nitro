package ast

import "github.com/dcaiafa/nitro/internal/token"

type Unit struct {
	Package token.Token
	Decls   ASTs
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		ctx.Analysis.SetPackage(u.Package.Str)
	}

	ctx.RunPassChild(u, u.Decls, pass)
}
