package ast

import "github.com/dcaiafa/nitro/internal/token"

type GoType struct {
	Ref     bool
	Package token.Token
	ID      token.Token
}

func (t *GoType) RunPass(ctx *Context, pass Pass) {}
