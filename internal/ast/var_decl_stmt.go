package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type VarDeclStmt struct {
	astBase
	Vars       []token.Token
	InitValues Exprs

	syms []symbol.Symbol
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if len(s.InitValues) != 0 {
			if len(s.Vars) != len(s.InitValues) {
				if funcCall, ok := s.InitValues[0].(*FuncCallExpr); ok && len(s.InitValues) == 1 {
					funcCall.RetN = len(s.Vars)
				} else {
					ctx.Failf(s.Pos(), "assigment mismatch: %v variables but %v values",
						len(s.Vars), len(s.InitValues))
					return
				}
			}
		}

	case Emit:
		emitter := ctx.Emitter()

		for _, sym := range s.syms {
			emitVariableInit(ctx, s.Pos(), sym)
		}

		if s.InitValues != nil {
			for _, sym := range s.syms {
				emitSymbolRefPush(s.Pos(), emitter, sym)
			}
		}
	}

	if s.InitValues != nil {
		ctx.RunPassChild(s, s.InitValues, pass)
	}

	switch pass {
	case Check:
		s.syms = make([]symbol.Symbol, len(s.Vars))
		for i, v := range s.Vars {
			s.syms[i] = AddVariable(ctx, v.Str, v.Pos)
			if s.syms[i] == nil {
				return
			}
		}

	case Emit:
		if s.InitValues != nil {
			emitter := ctx.Emitter()
			emitter.Emit(s.Pos(), vm.OpStore, uint32(len(s.syms)), 0)
		}
	}
}
