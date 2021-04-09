package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type YieldStmt struct {
	astBase
	Values Exprs
}

func (s *YieldStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		fn := ctx.CurrentFunc()
		if fn == nil {
			ctx.Failf(s.Pos(), "cannot yield outside of a function")
			return
		}
		fn.MarkEnumerable()

	case Emit:
		ctx.Emitter().Emit(s.Pos(), runtime.OpNewBool, 1, 0)
	}

	ctx.RunPassChild(s, s.Values, pass)

	if pass == Emit {
		ctx.Emitter().Emit(s.Pos(), runtime.OpYield, uint32(len(s.Values))+1, 0)
	}
}
