package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type VarDeclStmt struct {
	PosImpl
	Vars       []token.Token
	InitValues Exprs

	init *VarDeclInit
}

func (s *VarDeclStmt) RunPass(ctx *Context, pass Pass) {
	parentFn := ctx.CurrentFunc()

	switch pass {
	case Rewrite:
		s.init = &VarDeclInit{
			PosImpl:    s.PosImpl,
			InitValues: s.InitValues,
		}
		if parentFn == nil {
			ctx.Package().AddInitStmt(s.init)
		}

	case CreateGlobals:
		if parentFn == nil {
			scope := ctx.GetScope(scope.Package)
			s.init.Syms = make([]symbol.Symbol, len(s.Vars))
			for i, v := range s.Vars {
				g := ctx.Package().NewGlobal()
				g.SetName(v.Str)
				g.SetPos(v.Pos)
				if !scope.PutSymbol(ctx, g) {
					return
				}
				s.init.Syms[i] = g
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
	}

	if parentFn != nil {
		// If it is a global variable, then s.init is rooted at the package's $init.
		ctx.RunPassChild(s, s.init, pass)
	}

	switch pass {
	case Check:
		if parentFn != nil {
			scope := ctx.GetScope(scope.Block)
			s.init.Syms = make([]symbol.Symbol, len(s.Vars))
			for i, v := range s.Vars {
				l := parentFn.NewLocal()
				l.SetName(v.Str)
				l.SetPos(v.Pos)
				if !scope.PutSymbol(ctx, l) {
					return
				}
				s.init.Syms[i] = l
			}
		}
	}
}

type VarDeclInit struct {
	PosImpl

	Syms       []symbol.Symbol
	InitValues Exprs
}

func (v *VarDeclInit) RunPass(ctx *Context, pass Pass) {
	emitter := ctx.Emitter()

	if pass == Emit {
		for _, sym := range v.Syms {
			emitVariableInit(ctx, v.Pos(), sym)
		}

		if v.InitValues != nil {
			for _, sym := range v.Syms {
				emitSymbolRefPush(v.Pos(), emitter, sym)
			}
		}
	}

	ctx.RunPassChild(v, v.InitValues, pass)

	if pass == Emit {
		if v.InitValues != nil {
			emitter := ctx.Emitter()
			emitter.Emit(v.Pos(), vm.OpStore, uint32(len(v.Syms)), 0)
		}
	}
}
