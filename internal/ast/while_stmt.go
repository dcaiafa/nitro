package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type WhileStmt struct {
	astBase

	Predicate Expr
	Block     *StmtBlock
}

func (s *WhileStmt) RunPass(ctx *Context, pass Pass) {
	var emitter *runtime.Emitter
	var begin, end *runtime.Label

	if pass == Emit {
		emitter = ctx.Emitter()
		begin = emitter.NewLabel()
		end = emitter.NewLabel()
		emitter.ResolveLabel(begin)
	}

	ctx.RunPassChild(s, s.Predicate, pass)

	if pass == Emit {
		emitter.EmitJump(s.Predicate.Pos(), runtime.OpJumpIfFalse, end, 0)
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		emitter.EmitJump(s.Block.Pos(), runtime.OpJump, begin, 0)
		emitter.ResolveLabel(end)
	}
}
