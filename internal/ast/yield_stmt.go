package ast

import "github.com/dcaiafa/nitro/internal/vm"

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
		if fn.IterNRet() != 0 && fn.IterNRet() != len(s.Values) {
			ctx.Failf(
				s.Pos(),
				"all yield statements inside an iterator must return the same "+
					"number of values")
			return
		}
		fn.SetIterNRet(len(s.Values))

	case Emit:
		ctx.Emitter().Emit(s.Pos(), vm.OpNewBool, 1, 0)
	}

	ctx.RunPassChild(s, s.Values, pass)

	if pass == Emit {
		ctx.Emitter().Emit(s.Pos(), vm.OpIterYield, uint32(len(s.Values))+1, 0)
	}
}
