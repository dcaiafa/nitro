package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

func AddVariable(ctx *Context, name string, pos token.Pos) symbol.Symbol {
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
