package ast

import (
	"github.com/dcaiafa/nitro/internal/stub/analysis"
	"github.com/dcaiafa/nitro/internal/token"
)

type TypeRef struct {
	Iter    bool
	Nilable bool
	ID      token.Token

	Type analysis.Type
}

func (r *TypeRef) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		var ok bool
		r.Type, ok = ctx.Analysis.GetGoType(r.ID.Str)
		if !ok {
			ctx.Failf(token.Pos{}, "Undefined type %q", r.ID.Str)
			return
		}
		r.Type.Nilable = r.Nilable
	}
}
