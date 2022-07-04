package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type FuncStmt struct {
	Func

	Name string

	sym symbol.Symbol
}

func (s *FuncStmt) RunPass(ctx *Context, pass Pass) {
	parentFn := ctx.CurrentFunc()

	switch pass {
	case CreateGlobals:
		if parentFn == nil {
			s.sym = &symbol.FuncSymbol{}
			s.sym.SetReadOnly(true)
			s.sym.SetName(s.Name)
			s.sym.SetPos(s.Pos())

			if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
				return
			}
		}

	case Check:
		s.DebugName = s.Name

		if parentFn != nil {
			s.sym = parentFn.NewLocal()
			s.sym.SetName(s.Name)
			s.sym.SetPos(s.Pos())
			s.IsClosure = true

			if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
				return
			}
		}

	case Emit:
		if localSym, ok := s.sym.(*symbol.LocalVarSymbol); ok {
			emitVariableInit(ctx, s.Pos(), localSym)
			emitSymbolRefPush(s.Pos(), ctx.Emitter(), localSym)
		}
	}

	ctx.RunPassChild(s, &s.Func, pass)

	switch pass {
	case Check:
		if fnSym, ok := s.sym.(*symbol.FuncSymbol); ok {
			fnSym.IdxFunc = s.IdxFunc()
		}

	case Emit:
		if _, ok := s.sym.(*symbol.LocalVarSymbol); ok {
			// The prefix emitted the PushLocalRef. `Func` emitted the closure. Now
			// emit the `Store` to place the closure into the local var.
			ctx.Emitter().Emit(s.Pos(), vm.OpStore, 1, 0)
		}
	}
}
