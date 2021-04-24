package ast

import (
	"github.com/dcaiafa/nitro/internal/vm"
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
		} else if fn != nil && fn.IsIter() && len(s.Values) != 0 {
			ctx.Failf(s.Pos(), "'return' must return no values when inside an iterator")
			return
		}

		retOp := vm.OpRet
		if fn != nil && fn.IsIter() {
			retOp = vm.OpIterRet
		}

		ctx.Emitter().Emit(s.Pos(), retOp, uint32(len(s.Values)), 0)
	}
}
