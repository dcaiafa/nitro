package ast

import (
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/internal/token"
)

type MemberAccess struct {
	astBase
	Target Expr
	Member token.Token
}

func (a *MemberAccess) isExpr() {}

func (a *MemberAccess) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(a, a.Target, pass)

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.Emit(
			a.Pos(), vm.OpLoadLiteral,
			uint32(emitter.AddString(a.Member.Str)), 0)
		if _, isLValue := ctx.Parent().(*LValue); isLValue {
			emitter.Emit(a.Pos(), vm.OpObjectGetRef, 0, 0)
		} else {
			emitter.Emit(a.Pos(), vm.OpObjectGet, 0, 0)
		}
	}
}
