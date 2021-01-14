package ast

import (
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
)

func AddVariable(ctx *Context, name string, pos token.Pos) types.Symbol {
	fn := ctx.CurrentFunc()
	if fn != nil {
		l := fn.NewLocal()
		l.SetName(name)
		l.SetPos(pos)
		if !ctx.CurrentScope().PutSymbol(ctx, l) {
			return nil
		}
		return l
	}

	g := ctx.Main().NewGlobal()
	g.SetName(name)
	g.SetPos(pos)
	if !ctx.CurrentScope().PutSymbol(ctx, g) {
		return nil
	}

	return g
}
