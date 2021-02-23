package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type FuncStmt struct {
	Func

	Name string

	sym types.Symbol
}

func (s *FuncStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		parentFn := ctx.CurrentFunc()

		if parentFn != nil {
			s.sym = parentFn.NewLocal()
			s.Func.IsClosure = true
		} else {
			s.sym = &types.FuncSymbol{}
		}

		s.sym.SetName(s.Name)
		s.sym.SetPos(s.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
			return
		}

	case Emit:
		if localSym, ok := s.sym.(*types.LocalVarSymbol); ok {
			ctx.Emitter().Emit(runtime.OpLoadLocalRef, uint16(localSym.LocalNdx), 0)
		}
	}

	ctx.RunPassChild(s, &s.Func, pass)

	switch pass {
	case Check:
		if fnSym, ok := s.sym.(*types.FuncSymbol); ok {
			fnSym.FnNdx = s.Func.FnNdx
		}

	case Emit:
		if _, ok := s.sym.(*types.LocalVarSymbol); ok {
			// The prefix emitted the PushLocalRef. `Func` emitted the closure. Now
			// emit the `Store` to place the closure into the local var.
			ctx.Emitter().Emit(runtime.OpStore, 1, 0)
		}
	}
}
