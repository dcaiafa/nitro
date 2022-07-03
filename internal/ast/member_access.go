package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type MemberAccess struct {
	PosImpl
	Target Expr
	Member token.Token

	ModuleMember symbol.Symbol
}

func (a *MemberAccess) isExpr() {}

func (a *MemberAccess) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(a, a.Target, pass)

	switch pass {
	case Check:
		simpleRef, ok := a.Target.(*SimpleRef)
		if ok {
			if simpleRef.ModuleRef != nil {
				a.ModuleMember = simpleRef.ModuleRef.Module.Scope.GetSymbol(a.Member.Str)
				if a.ModuleMember == nil {
					ctx.Failf(a.Pos(), "module %v does not have symbol %v",
						simpleRef.ModuleRef.Module.Name,
						a.Member.Str)
					return
				}
			}
		}

	case Emit:
		emitter := ctx.Emitter()
		if a.ModuleMember == nil {
			emitter.Emit(
				a.Pos(), vm.OpLoadLiteral,
				uint32(emitter.AddString(a.Member.Str)), 0)
			if _, isLValue := ctx.Parent().(*LValue); isLValue {
				emitter.Emit(a.Pos(), vm.OpObjectGetRef, 0, 0)
			} else {
				emitter.Emit(a.Pos(), vm.OpObjectGet, 0, 0)
			}
		} else {
			if _, isLValue := ctx.Parent().(*LValue); isLValue {
				ctx.Failf(a.Pos(), "cannot assign to module")
				return
			}
			emitSymbolPush(a.Pos(), ctx.Emitter(), a.ModuleMember)
		}
	}
}
