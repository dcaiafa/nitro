package ast

import "github.com/dcaiafa/nitro/internal/token"

type ConstDecl struct {
	Doc     *Doc
	ID      token.Token
	TypeRef *TypeRef
	Value   *ConstValue
}

func (d *ConstDecl) RunPass(ctx *Context, pass Pass) {
}
