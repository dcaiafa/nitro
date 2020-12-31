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

	Sym *types.LocalVarSymbol
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		symName := s.VarName.Str

		s.Sym = &types.LocalVarSymbol{}
		s.Sym.SetName(symName)
		s.Sym.SetPos(s.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, s.Sym) {
			return
		}

		fn := ctx.CurrentFunc().Sym
		fn.Locals = append(fn.Locals, s.Sym)

	case Emit:
		if s.InitValue != nil {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpPushLocalRef, uint16(s.Sym.Local), 0)
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
			emitter.Emit(runtime.OpStore, uint16(s.Sym.Local), 0)
		}
	}
}
