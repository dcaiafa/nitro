package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type SimpleRef struct {
	astBase
	ID token.Token

	sym       symbol.Symbol
	ModuleRef *symbol.ModuleRef
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

		var ok bool
		r.ModuleRef, ok = r.sym.(*symbol.ModuleRef)
		if ok {
			if _, ok := ctx.Parent().(*MemberAccess); !ok {
				ctx.Failf(
					r.Pos(),
					"%v is a module reference, and cannot be used as a value",
					r.ID.Str)
				return
			}
		}

	case Emit:
		if r.ModuleRef == nil {
			emit := emitSymbolPush
			_, isLValue := ctx.Parent().(*LValue)
			if isLValue {
				emit = emitSymbolRefPush
			}
			emit(r.Pos(), ctx.Emitter(), r.sym)
		}
	}
}
