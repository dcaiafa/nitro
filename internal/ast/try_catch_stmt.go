package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type TryCatchStmt struct {
	PosImpl

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
	var beginCatch, endCatch *vm.Label

	if pass == Emit {
		e := ctx.Emitter()
		beginCatch = e.NewLabel()
		endCatch = e.NewLabel()
		e.EmitJump(s.Pos(), vm.OpBeginTry, beginCatch, 0)
	}

	ctx.RunPassChild(s, s.tryBlock, pass)

	if pass == Emit {
		ctx.Emitter().EmitJump(s.Pos(), vm.OpEndTry, endCatch, 0)
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
	PosImpl

	catchVar *token.Token
	stmts    AST

	scope    *symbol.Scope
	catchSym symbol.Symbol
}

func (b *catchBlock) Scope() *symbol.Scope {
	return b.scope
}

func (b *catchBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = symbol.NewScope()

		if b.catchVar != nil {
			l := ctx.CurrentFunc().NewLocal()
			l.SetName(b.catchVar.Str)
			l.SetPos(b.catchVar.Pos)
			if !b.scope.PutSymbol(ctx, l) {
				return
			}
			b.catchSym = l
		}

	case Emit:
		if b.catchSym != nil {
			emitVariableInit(ctx, b.stmts.Pos(), b.catchSym)
			emitSymbolRefPush(b.stmts.Pos(), ctx.Emitter(), b.catchSym)
			ctx.Emitter().Emit(b.stmts.Pos(), vm.OpSwap, 1, 0)
			ctx.Emitter().Emit(b.stmts.Pos(), vm.OpStore, 1, 0)
		} else {
			ctx.Emitter().Emit(b.stmts.Pos(), vm.OpPop, 1, 0)
		}
	}

	ctx.RunPassChild(b, b.stmts, pass)
}
