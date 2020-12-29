package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
)

type ReturnStmt struct {
	astBase
	Values Exprs
}

func (s *ReturnStmt) RunPass(ctx *Context, pass Pass) {
	ctx.Push(s)
	s.Values.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		ctx.Emitter().Emit(runtime.OpRet, uint64(len(s.Values)))
	}
}
