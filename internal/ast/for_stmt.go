package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type ForStmt struct {
	astBase
	ForVars  ASTs
	IterExpr Expr
	Block    AST

	scope *symbol.Scope
	iter  symbol.Symbol
	begin *runtime.Label
	end   *runtime.Label
}

func (s *ForStmt) Scope() *symbol.Scope {
	return s.scope
}

func (s *ForStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		s.scope = symbol.NewScope()
		s.iter = AddVariableToScope(ctx, s.scope, "$iter", s.IterExpr.Pos())
	}

	ctx.RunPassChild(s, s.ForVars, pass)

	if pass == Emit {
		emitSymbolRefPush(s.Pos(), ctx.Emitter(), s.iter)
	}

	ctx.RunPassChild(s, s.IterExpr, pass)

	if pass == Emit {
		e := ctx.Emitter()
		s.begin = e.NewLabel()
		s.end = e.NewLabel()
		e.Emit(s.Pos(), runtime.OpMakeIter, 0, 0)
		e.Emit(s.Pos(), runtime.OpStore, 1, 0)
		e.ResolveLabel(s.begin)
		for _, v := range s.ForVars {
			emitSymbolRefPush(s.Pos(), e, v.(*ForVar).sym)
		}
		emitSymbolPush(s.Pos(), e, s.iter)
		e.Emit(s.Pos(), runtime.OpCall, 0, uint16(len(s.ForVars)+1))
		e.EmitJump(s.Pos(), runtime.OpNext, s.end, uint16(len(s.ForVars)+1))
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		e := ctx.Emitter()
		e.EmitJump(s.Pos(), runtime.OpJump, s.begin, 0)
		e.ResolveLabel(s.end)
	}
}

func (s *ForStmt) EmitBreak(pos token.Pos, e *runtime.Emitter) {
	e.EmitJump(pos, runtime.OpJump, s.end, 0)
}

func (s *ForStmt) EmitContinue(pos token.Pos, e *runtime.Emitter) {
	e.EmitJump(pos, runtime.OpJump, s.begin, 0)
}

type ForVar struct {
	astBase
	VarName token.Token

	sym symbol.Symbol
}

func (s *ForVar) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		s.sym = AddVariable(ctx, s.VarName.Str, s.VarName.Pos)
		if s.sym == nil {
			return
		}
	}
}
