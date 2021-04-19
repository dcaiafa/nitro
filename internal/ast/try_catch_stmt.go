package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type TryCatchStmt struct {
	astBase

	tryBlock   AST
	catchBlock *catchBlock
}

func NewTryCatchStmt(
	tryBlock AST,
	catchVar *token.Token,
	catchStmts AST,
) *TryCatchStmt {
	return &TryCatchStmt{
		tryBlock: tryBlock,
		catchBlock: &catchBlock{
			catchVar: catchVar,
			stmts:    catchStmts,
		},
	}
}

func (s *TryCatchStmt) RunPass(ctx *Context, pass Pass) {
	var beginCatch, endCatch *runtime.Label

	if pass == Emit {
		e := ctx.Emitter()
		beginCatch = e.NewLabel()
		endCatch = e.NewLabel()
		e.EmitJump(s.Pos(), runtime.OpBeginTry, beginCatch, 0)
	}

	ctx.RunPassChild(s, s.tryBlock, pass)

	if pass == Emit {
		ctx.Emitter().EmitJump(s.Pos(), runtime.OpEndTry, endCatch, 0)
	}

	if pass == Emit {
		ctx.Emitter().ResolveLabel(beginCatch)
	}

	ctx.RunPassChild(s, s.catchBlock, pass)

	if pass == Emit {
		ctx.Emitter().ResolveLabel(endCatch)
	}
}

type catchBlock struct {
	astBase

	catchVar *token.Token
	stmts    AST

	scope    *symbol.Scope
	catchSym symbol.Symbol
}

func (b *catchBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = symbol.NewScope()

		if b.catchVar != nil {
			b.catchSym = AddVariable(ctx, b.catchVar.Str, b.catchVar.Pos)
			if b.catchSym == nil {
				return
			}
		}

	case Emit:
		if b.catchSym != nil {
			emitSymbolRefPush(b.stmts.Pos(), ctx.Emitter(), b.catchSym)
			ctx.Emitter().Emit(b.stmts.Pos(), runtime.OpSwap, 1, 0)
			ctx.Emitter().Emit(b.stmts.Pos(), runtime.OpStore, 1, 0)
		} else {
			ctx.Emitter().Emit(b.stmts.Pos(), runtime.OpPop, 1, 0)
		}
	}

	ctx.RunPassChild(b, b.stmts, pass)
}

func (b *catchBlock) Scope() *symbol.Scope {
	return b.scope
}
