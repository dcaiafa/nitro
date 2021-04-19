package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type FuncStmt struct {
	Func

	Name string

	sym symbol.Symbol
}

func (s *FuncStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		parentFn := ctx.CurrentFunc()

		if parentFn != nil {
			s.sym = parentFn.NewLocal()
			s.Func.IsClosure = true
		} else {
			s.sym = &symbol.FuncSymbol{}
		}

		s.Func.DebugName = s.Name
		s.sym.SetName(s.Name)
		s.sym.SetPos(s.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
			return
		}

	case Emit:
		if localSym, ok := s.sym.(*symbol.LocalVarSymbol); ok {
			ctx.Emitter().Emit(s.Pos(), runtime.OpLoadLocalRef, uint32(localSym.LocalNdx), 0)
		}
	}

	ctx.RunPassChild(s, &s.Func, pass)

	switch pass {
	case Check:
		if fnSym, ok := s.sym.(*symbol.FuncSymbol); ok {
			fnSym.IdxFunc = s.Func.IdxFunc()
		}

	case Emit:
		if _, ok := s.sym.(*symbol.LocalVarSymbol); ok {
			// The prefix emitted the PushLocalRef. `Func` emitted the closure. Now
			// emit the `Store` to place the closure into the local var.
			ctx.Emitter().Emit(s.Pos(), runtime.OpStore, 1, 0)
		}
	}
}
