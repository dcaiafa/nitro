package ast

import "github.com/dcaiafa/nitro/internal/token"

type ExecHome struct {
	PosImpl
	suggar AST
}

func (e *ExecHome) isExpr() {}

func (e *ExecHome) RunPass(ctx *Context, pass Pass) {
	if pass == Rewrite {
		e.suggar = &FuncCallExpr{
			Target: &SimpleRef{
				ID: token.Token{
					Type: token.String,
					Str:  "$home",
				},
			},
			RetN: 1,
		}
	}
	ctx.RunPassChild(e, e.suggar, pass)
}
