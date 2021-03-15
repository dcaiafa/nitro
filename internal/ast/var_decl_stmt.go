package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
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
		s.syms = make([]symbol.Symbol, len(s.Vars))
		for i, v := range s.Vars {
			s.syms[i] = AddVariable(ctx, v.Str, v.Pos)
			if s.syms[i] == nil {
				return
			}
		}
		if len(s.InitValues) != 0 {
			if len(s.Vars) != len(s.InitValues) {
				if funcCall, ok := s.InitValues[0].(*FuncCallExpr); ok && len(s.InitValues) == 1 {
					funcCall.RetN = len(s.InitValues)
				} else {
					ctx.Failf(s.Pos(), "assigment mismatch: %v variables but %v values",
						len(s.Vars), len(s.InitValues))
				}
			}
		}

	case Emit:
		if s.InitValues != nil {
			emitter := ctx.Emitter()
			for _, sym := range s.syms {
				switch sym := sym.(type) {
				case *symbol.LocalVarSymbol:
					emitter.Emit(s.Pos(), runtime.OpLoadLocalRef, uint16(sym.LocalNdx), 0)
				case *symbol.GlobalVarSymbol:
					emitter.Emit(s.Pos(), runtime.OpLoadGlobalRef, uint16(sym.GlobalNdx), 0)
				}
			}
		}
	}

	if s.InitValues != nil {
		ctx.RunPassChild(s, s.InitValues, pass)
	}

	switch pass {
	case Emit:
		if s.InitValues != nil {
			emitter := ctx.Emitter()
			emitter.Emit(s.Pos(), runtime.OpStore, uint16(len(s.syms)), 0)
		}
	}
}
