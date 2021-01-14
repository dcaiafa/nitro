package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
)

type VarDeclStmt struct {
	astBase
	VarName   token.Token
	InitValue Expr

	Sym types.Symbol
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		s.Sym = AddVariable(ctx, s.VarName.Str, s.Pos())
		if s.Sym == nil {
			return
		}

	case Emit:
		if s.InitValue != nil {
			emitter := ctx.Emitter()

			switch sym := s.Sym.(type) {
			case *types.LocalVarSymbol:
				emitter.Emit(runtime.OpPushLocalRef, uint16(sym.LocalNdx), 0)

			case *types.GlobalVarSymbol:
				emitter.Emit(runtime.OpPushGlobalRef, uint16(sym.GlobalNdx), 0)
			}
		}
	}

	if s.InitValue != nil {
		ctx.Push(s)
		s.InitValue.RunPass(ctx, pass)
		ctx.Pop()
	}

	switch pass {
	case Emit:
		if s.InitValue != nil {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpStore, 0, 0)
		}
	}
}
