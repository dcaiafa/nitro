package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type ForStmt struct {
	PosImpl
	ForVars  ASTs
	IterExpr Expr
	Block    AST

	scope scope.Scope
	iter  symbol.Symbol
	begin *vm.Label
	end   *vm.Label
}

var _ Scope = (*ForStmt)(nil)

func (s *ForStmt) IsRepeatableScope() {}

func (s *ForStmt) Scope() scope.Scope {
	return s.scope
}

func (s *ForStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		s.scope = scope.NewScope(scope.Block)
		l := ctx.CurrentFunc().NewLocal()
		l.SetName("$iter")
		l.SetPos(s.IterExpr.Pos())
		if !s.scope.PutSymbol(ctx, l) {
			return
		}
		s.iter = l
	}

	ctx.RunPassChild(s, s.ForVars, pass)

	if pass == Emit {
		emitVariableInit(ctx, s.Pos(), s.iter)
		emitSymbolRefPush(s.Pos(), ctx.Emitter(), s.iter)
	}

	ctx.RunPassChild(s, s.IterExpr, pass)

	if pass == Emit {
		e := ctx.Emitter()
		s.begin = e.NewLabel()
		s.end = e.NewLabel()

		// Make iterator and store it in "$iter".
		e.Emit(s.Pos(), vm.OpMakeIter, 0, 0)
		e.Emit(s.Pos(), vm.OpStore, 1, 0)

		e.ResolveLabel(s.begin)

		for _, v := range s.ForVars {
			emitSymbolRefPush(s.Pos(), e, v.(*ForVar).sym)
		}
		emitSymbolPush(s.Pos(), e, s.iter)
		e.EmitJump(s.Pos(), vm.OpNext, s.end, uint16(len(s.ForVars)))
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		e := ctx.Emitter()
		e.EmitJump(s.Pos(), vm.OpJump, s.begin, 0)
		e.ResolveLabel(s.end)
	}
}

func (s *ForStmt) EmitBreak(pos token.Pos, e *vm.Emitter) {
	e.EmitJump(pos, vm.OpJump, s.end, 0)
}

func (s *ForStmt) EmitContinue(pos token.Pos, e *vm.Emitter) {
	e.EmitJump(pos, vm.OpJump, s.begin, 0)
}

type ForVar struct {
	PosImpl
	VarName token.Token

	sym symbol.Symbol
}

func (s *ForVar) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		l := ctx.CurrentFunc().NewLocal()
		l.SetName(s.VarName.Str)
		l.SetPos(s.VarName.Pos)
		if !ctx.GetScope(scope.Block).PutSymbol(ctx, l) {
			return
		}
		s.sym = l

	case Emit:
		emitVariableInit(ctx, s.Pos(), s.sym)
	}
}
