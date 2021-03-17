package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
)

type TryCatchStmt struct {
	astBase

	TryBlock   AST
	CatchVar   *token.Token
	CatchBlock AST

	catchVar symbol.Symbol
}

func (s *TryCatchStmt) RunPass(ctx *Context, pass Pass) {
	var beginCatch, endCatch *runtime.Label
	switch pass {
	case Emit:
		e := ctx.Emitter()
		beginCatch = e.NewLabel()
		endCatch = e.NewLabel()
		e.EmitJump(s.Pos(), runtime.OpStartTry, beginCatch)
	}

	ctx.RunPassChild(s, s.TryBlock, pass)

	switch pass {
	case Emit:
		ctx.Emitter().EmitJump(s.Pos(), runtime.OpEndTry, endCatch)
	}

	switch pass {
	case Check:
		if s.CatchVar != nil {
			ctx.Push(s.CatchBlock)
			s.catchVar = AddVariable(ctx, s.CatchVar.Str, s.CatchVar.Pos)
			ctx.Pop()

			if s.catchVar == nil {
				return
			}
		}
	case Emit:
		ctx.Emitter().ResolveLabel(beginCatch)
	}

	ctx.RunPassChild(s, s.CatchBlock, pass)

	switch pass {
	case Emit:
		ctx.Emitter().ResolveLabel(endCatch)
	}
}
