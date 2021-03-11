package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type IndexExpr struct {
	astBase
	Target Expr
	Index  Expr
}

func (e *IndexExpr) isExpr() {}

func (e *IndexExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Target, pass)
	ctx.RunPassChild(e, e.Index, pass)

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		if _, isLValue := ctx.Parent().(*LValue); isLValue {
			emitter.Emit(runtime.OpObjectGetRef, 0, 0)
		} else {
			emitter.Emit(runtime.OpObjectGet, 0, 0)
		}
	}
}
