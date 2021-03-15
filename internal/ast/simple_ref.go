package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type SimpleRef struct {
	astBase
	ID token.Token

	sym symbol.Symbol
}

func (r *SimpleRef) isExpr() {}

func (r *SimpleRef) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		symName := r.ID.Str

		r.sym = ctx.FindSymbol(symName)
		if r.sym == nil {
			ctx.Failf(r.Pos(), "Symbol %q not found.", symName)
			return
		}

	case Emit:
		emit := emitSymbolPush
		_, isLValue := ctx.Parent().(*LValue)
		if isLValue {
			emit = emitSymbolRefPush
		}
		emit(r.Pos(), ctx.Emitter(), r.sym)
	}
}
