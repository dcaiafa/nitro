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
		if fn.IteratorNRet() != 0 && fn.IteratorNRet() != len(s.Values) {
			ctx.Failf(
				s.Pos(),
				"all yield statements inside an iterator must return the same "+
					"number of values")
			return
		}
		fn.SetIteratorNRet(len(s.Values))

	case Emit:
		ctx.Emitter().Emit(s.Pos(), runtime.OpNewBool, 1, 0)
	}

	ctx.RunPassChild(s, s.Values, pass)

	if pass == Emit {
		ctx.Emitter().Emit(s.Pos(), runtime.OpIterYield, uint32(len(s.Values))+1, 0)
	}
}
