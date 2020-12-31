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

	sym *types.LocalVarSymbol
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateAndResolveNames:
		symName := s.VarName.Str

		owner := ctx.CurrentFunc()
		s.sym = &types.LocalVarSymbol{}
		s.sym.SetName(symName)
		s.sym.SetOwner(owner)
		s.sym.SetPos(s.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
			return
		}
		owner.Locals = append(owner.Locals, s.sym)

	case Emit:
		if s.InitValue != nil {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpPushLocalRef, uint16(s.sym.Local), 0)
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
			emitter.Emit(runtime.OpStore, uint16(s.sym.Local), 0)
		}
	}
}
