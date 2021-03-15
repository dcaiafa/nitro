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
}

func (s *ForStmt) Scope() *symbol.Scope {
	return s.scope
}

func (s *ForStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Check {
		s.scope = symbol.NewScope()
		s.iter = AddVariable(ctx, "$iter", s.IterExpr.Pos())
	}

	ctx.RunPassChild(s, s.ForVars, pass)

	if pass == Emit {
		emitSymbolRefPush(s.Pos(), ctx.Emitter(), s.iter)
	}

	ctx.RunPassChild(s, s.IterExpr, pass)

	var begin, end *runtime.Label
	if pass == Emit {
		e := ctx.Emitter()
		begin = e.NewLabel()
		end = e.NewLabel()
		e.Emit(s.Pos(), runtime.OpMakeIter, 0, 0)
		e.Emit(s.Pos(), runtime.OpStore, 1, 0)
		e.ResolveLabel(begin)
		for _, v := range s.ForVars {
			emitSymbolRefPush(s.Pos(), e, v.(*ForVar).sym)
		}
		emitSymbolPush(s.Pos(), e, s.iter)
		e.Emit(s.Pos(), runtime.OpCall, 0, byte(len(s.ForVars)+1))
		e.EmitJump(s.Pos(), runtime.OpJumpIfFalse, end)
		e.Emit(s.Pos(), runtime.OpStore, uint16(len(s.ForVars)), 0)
	}

	ctx.RunPassChild(s, s.Block, pass)

	if pass == Emit {
		e := ctx.Emitter()
		e.EmitJump(s.Pos(), runtime.OpJump, begin)
		e.ResolveLabel(end)
		e.Emit(s.Pos(), runtime.OpPop, uint16(len(s.ForVars)*2), 0)
	}
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
