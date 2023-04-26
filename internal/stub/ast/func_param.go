package ast

import "github.com/dcaiafa/nitro/internal/token"

type FuncParam struct {
	ID           token.Token
	Type         *TypeRef
	DefaultValue *ConstValue
	VarArg       bool
}

func (p *FuncParam) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(p, p.Type, pass)
	if p.DefaultValue != nil {
		ctx.RunPassChild(p, p.DefaultValue, pass)
	}
}
