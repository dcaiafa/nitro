package ast

import (
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
	if s.InitValue != nil {
		ctx.Push(s)
		s.InitValue.RunPass(ctx, pass)
		ctx.Pop()
	}

	owner := ctx.CurrentFunc()

	switch pass {
	case CreateAndResolveNames:
		symName := s.VarName.Str

		currentScope := ctx.CurrentScope()
		existingSym := currentScope.GetSymbol(symName)
		if existingSym != nil {
			ctx.Failf(
				s.Pos(),
				"There is already something named %q in the current scope.",
				symName)
			ctx.Failf(s.Pos(), "%q was previously declared here.", symName)
			return
		}

		s.sym = &types.LocalVarSymbol{}
		s.sym.SetName(symName)
		s.sym.SetOwner(owner)
		s.sym.SetPos(s.Pos())
		currentScope.PutSymbol(s.sym)
		owner.Locals = append(owner.Locals, s.sym)

	case Emit:
	}
}
