package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Unit struct {
	astBase

	Meta    ASTs
	Imports ASTs
	Block   *StmtBlock

	scope *symbol.Scope
	fn    int
}

func (m *Unit) Scope() *symbol.Scope {
	return m.scope
}

func (m *Unit) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		m.scope = symbol.NewScope()
		m.fn = ctx.Emitter().NewFn("main")

	case Emit:
		emitter := ctx.Emitter()
		emitter.PushFn(m.fn)
		emitter.Emit(m.Pos(), vm.OpInitCallFrame, 0, 0)
	}

	ctx.RunPassChild(m, m.Meta, pass)
	ctx.RunPassChild(m, m.Imports, pass)
	ctx.RunPassChild(m, m.Block, pass)

	switch pass {
	case Emit:
		synthesizeReturn := len(m.Block.Stmts) == 0
		if !synthesizeReturn {
			_, endsWithReturn := m.Block.Stmts[len(m.Block.Stmts)-1].(*ReturnStmt)
			synthesizeReturn = !endsWithReturn
		}
		if synthesizeReturn {
			ctx.Emitter().Emit(m.Pos(), vm.OpRet, 0, 0)
		}

		ctx.Emitter().PopFn()
	}
}
