package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type VarDeclStmt struct {
	PosImpl
	Vars       []token.Token
	InitValues Exprs

	syms []symbol.Symbol
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	parentFn := ctx.CurrentFunc()

	switch pass {
	case CreateGlobals:
		if parentFn == nil {
			scope := ctx.CurrentScope()
			s.syms = make([]symbol.Symbol, len(s.Vars))
			for i, v := range s.Vars {
				g := ctx.Main().NewGlobal()
				g.SetName(v.Str)
				g.SetPos(v.Pos)
				if !scope.PutSymbol(ctx, g) {
					return
				}
				s.syms[i] = g
			}
		}

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

		if parentFn == nil {
			emitter.PushFn(ctx.Main().initSym.IdxFunc)
			defer emitter.PopFn()
		}

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
		if parentFn != nil {
			scope := ctx.CurrentScope()
			s.syms = make([]symbol.Symbol, len(s.Vars))
			for i, v := range s.Vars {
				l := parentFn.NewLocal()
				l.SetName(v.Str)
				l.SetPos(v.Pos)
				if !scope.PutSymbol(ctx, l) {
					return
				}
				s.syms[i] = l
			}
		}

	case Emit:
		if s.InitValues != nil {
			emitter := ctx.Emitter()
			emitter.Emit(s.Pos(), vm.OpStore, uint32(len(s.syms)), 0)
		}
	}
}
