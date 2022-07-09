package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type SimpleRef struct {
	PosImpl
	ID token.Token

	sym symbol.Symbol

	Package *symbol.Package
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
		r.Package, ok = r.sym.(*symbol.Package)
		if ok {
			if _, ok := ctx.Parent().(*MemberAccess); !ok {
				ctx.Failf(
					r.Pos(),
					"%v is a package, and cannot be used as a value",
					r.ID.Str)
				return
			}
		}

		if r.sym.ReadOnly() {
			_, isLValue := ctx.Parent().(*LValue)
			if isLValue {
				ctx.Failf(
					r.Pos(),
					"%v is read-only and cannot be assigned to",
					r.ID.Str)
				return
			}
		}

	case Emit:
		if r.Package == nil {
			emit := emitSymbolPush
			_, isLValue := ctx.Parent().(*LValue)
			if isLValue {
				emit = emitSymbolRefPush
			}
			emit(r.Pos(), ctx.Emitter(), r.sym)
		}
	}
}
