package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Module struct {
	astBase

	Stmts ASTs

	scope *types.Scope
	fn    int
}

func (m *Module) Scope() *types.Scope {
	return m.scope
}

func (m *Module) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		m.scope = types.NewScope()
		m.fn = ctx.Emitter().NewFn()

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(m.fn)
		emitter.Emit(runtime.OpInitCallFrame, 0, 0)
	}

	ctx.Push(m)
	m.Stmts.RunPass(ctx, pass)
	ctx.Pop()

	switch pass {
	case Emit:
		synthesizeReturn := len(m.Stmts) == 0
		if !synthesizeReturn {
			_, endsWithReturn := m.Stmts[len(m.Stmts)-1].(*ReturnStmt)
			synthesizeReturn = !endsWithReturn
		}
		if synthesizeReturn {
			ctx.Emitter().Emit(runtime.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}
