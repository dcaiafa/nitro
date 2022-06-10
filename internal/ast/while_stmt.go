package ast

import (
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type WhileStmt struct {
	astBase

	Predicate Expr
	Block     *StmtBlock

	begin *vm.Label
	end   *vm.Label
}

func (s *WhileStmt) IsLiftableScope()   {}
func (s *WhileStmt) IsRepeatableScope() {}

func (s *WhileStmt) RunPass(ctx *Context, pass Pass) {
	var emitter *vm.Emitter

	if pass == Emit {
		emitter = ctx.Emitter()
		s.begin = emitter.NewLabel()
		s.end = emitter.NewLabel()
		emitter.ResolveLabel(s.begin)
	}

	ctx.RunPassChild(s, s.Predicate, pass)

	if pass == Emit {
		emitter.EmitJump(s.Predicate.Pos(), vm.OpJumpIfFalse, s.end, 0)
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		emitter.EmitJump(s.Block.Pos(), vm.OpJump, s.begin, 0)
		emitter.ResolveLabel(s.end)
	}
}

func (s *WhileStmt) EmitBreak(pos token.Pos, e *vm.Emitter) {
	e.EmitJump(pos, vm.OpJump, s.end, 0)
}

func (s *WhileStmt) EmitContinue(pos token.Pos, e *vm.Emitter) {
	e.EmitJump(pos, vm.OpJump, s.begin, 0)
}
