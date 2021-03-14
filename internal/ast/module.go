package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Module struct {
	astBase

	Block *StmtBlock

	scope *symbol.Scope
	fn    int
}

func (m *Module) Scope() *symbol.Scope {
	return m.scope
}

func (m *Module) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		m.scope = symbol.NewScope()
		m.fn = ctx.Emitter().NewFn()

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(m.fn)
		emitter.Emit(runtime.OpInitCallFrame, 0, 0)
	}

	ctx.RunPassChild(m, m.Block, pass)

	switch pass {
	case Emit:
		synthesizeReturn := len(m.Block.Stmts) == 0
		if !synthesizeReturn {
			_, endsWithReturn := m.Block.Stmts[len(m.Block.Stmts)-1].(*ReturnStmt)
			synthesizeReturn = !endsWithReturn
		}
		if synthesizeReturn {
			ctx.Emitter().Emit(runtime.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}
