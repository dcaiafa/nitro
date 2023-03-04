package ast

import (
	"path"

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
		name := i.Alias
		if name == "" {
			name = path.Base(i.ExpandedPackage)
		}

		pkg, index := ctx.GetDep(i.ExpandedPackage)
		imp := symbol.NewImport(pkg, index)
		imp.SetName(name)
		imp.SetPos(i.Pos())
		imp.SetReadOnly(true)

		unitScope := ctx.GetScope(scope.Unit)
		unitScope.PutSymbol(ctx, imp)
	}
}
