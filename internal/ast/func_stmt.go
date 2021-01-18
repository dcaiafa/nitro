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
		} else {
			s.sym = &types.FuncSymbol{}
		}

		s.sym.SetName(s.Name)
		s.sym.SetPos(s.Pos())
		if !ctx.CurrentScope().PutSymbol(ctx, s.sym) {
			return
		}
	}

	ctx.RunPassChild(s, &s.Func, pass)

	switch pass {
	case Check:
		if fnSym, ok := s.sym.(*types.FuncSymbol); ok {
			fnSym.FnNdx = s.Func.FnNdx
		}

	case Emit:
		if localSym, ok := s.sym.(*types.LocalVarSymbol); ok {
			emitter := ctx.Emitter()
			emitter.Emit(runtime.OpPushLocalRef, uint16(localSym.LocalNdx), 0)
			emitter.Emit(runtime.OpMakeClosure, uint16(s.Func.FnNdx), 0)
			emitter.Emit(runtime.OpStore, 0, 0)
		}
	}
}
