package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
)

type ReturnStmt struct {
	astBase
	Values Exprs
}

func (s *ReturnStmt) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.Values, pass)

	switch pass {
	case Emit:
		fn := ctx.CurrentFunc()
		if fn == nil && len(s.Values) != 0 {
			ctx.Failf(s.Pos(), "cannot return values outside of function")
			return
		} else if fn != nil && fn.IsIterator() && len(s.Values) != 0 {
			ctx.Failf(s.Pos(), "'return' must provide no values in an iterator")
			return
		}

		retOp := runtime.OpRet
		if fn != nil && fn.IsIterator() {
			retOp = runtime.OpIterRet
		}

		ctx.Emitter().Emit(s.Pos(), retOp, uint32(len(s.Values)), 0)
	}
}
