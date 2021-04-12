package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
)

type WhileStmt struct {
	astBase

	Predicate Expr
	Block     *StmtBlock

	begin *runtime.Label
	end   *runtime.Label
}

func (s *WhileStmt) RunPass(ctx *Context, pass Pass) {
	var emitter *runtime.Emitter

	if pass == Emit {
		emitter = ctx.Emitter()
		s.begin = emitter.NewLabel()
		s.end = emitter.NewLabel()
		emitter.ResolveLabel(s.begin)
	}

	ctx.RunPassChild(s, s.Predicate, pass)

	if pass == Emit {
		emitter.EmitJump(s.Predicate.Pos(), runtime.OpJumpIfFalse, s.end, 0)
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		emitter.EmitJump(s.Block.Pos(), runtime.OpJump, s.begin, 0)
		emitter.ResolveLabel(s.end)
	}
}

func (s *WhileStmt) EmitBreak(pos token.Pos, e *runtime.Emitter) {
	e.EmitJump(pos, runtime.OpJump, s.end, 0)
}

func (s *WhileStmt) EmitContinue(pos token.Pos, e *runtime.Emitter) {
	e.EmitJump(pos, runtime.OpJump, s.begin, 0)
}
