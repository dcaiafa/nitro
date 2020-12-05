package ast

import (
	"github.com/dcaiafa/nitro/internal/context"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
)

type SimpleRef struct {
	astBase
	ID token.Token

	sym *types.Symbol
}

func (s *SimpleRef) isExpr() {}

func (s *SimpleRef) RunPass(ctx *context.Context, pass context.Pass) {
	if pass == context.CreateAndResolveNames {
		symName := s.ID.Str
		s.sym, _ = ctx.Scopes().FindSymbol(symName)
		if s.sym == nil {
			ctx.Failf(s.Pos(), "Symbol %q not found.", symName)
		}
	}
}
