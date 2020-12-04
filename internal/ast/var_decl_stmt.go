package ast

import (
	"github.com/dcaiafa/nitro/internal/context"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/typecheck"
)

type VarDeclStmt struct {
	astBase
	VarName   token.Token
	InitValue Expr
}

func (s *VarDeclStmt) RunPass(ctx *context.Context, pass context.Pass) {
	s.runPassChildren(ctx, pass)

	// Create the symbol *after* running the pass on the right side to prevent the
	// right side from being able to reference it. E.g. var a = a + 1.
	if pass == context.CreateAndResolveNames {
		symName := s.VarName.Str

		existingSym := ctx.Scopes().Current().GetSymbol(symName)
		if existingSym != nil {
			ctx.Failf(
				s.VarName.Pos,
				"There is already something named %q in the current scope. "+
					"Declared right here: %v",
				symName, existingSym.Pos)
			return
		}

		sym := typecheck.NewSymbol(symName)
		sym.Pos = s.VarName.Pos
		ctx.Scopes().Current().PutSymbol(sym)
	}
}

func (s *VarDeclStmt) runPassChildren(ctx *context.Context, pass context.Pass) {
	if s.InitValue != nil {
		s.InitValue.RunPass(ctx, pass)
	}
}
