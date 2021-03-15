package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
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
			a.Pos(), runtime.OpLoadLiteral,
			uint16(emitter.AddString(a.Member.Str)), 0)
		if _, isLValue := ctx.Parent().(*LValue); isLValue {
			emitter.Emit(a.Pos(), runtime.OpObjectGetRef, 0, 0)
		} else {
			emitter.Emit(a.Pos(), runtime.OpObjectGet, 0, 0)
		}
	}
}
