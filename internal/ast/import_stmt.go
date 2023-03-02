package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Import struct {
	PosImpl

	Alias           string
	Package         string
	ExpandedPackage string
}

func (i *Import) RunPass(ctx *Context, pass Pass) {
	if pass == CreateGlobals {
		scope := ctx.GetScope(scope.Unit)
		global := new(symbol.GlobalVarSymbol)
	}
}
