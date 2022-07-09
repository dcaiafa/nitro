package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type ObjectLiteral struct {
	PosImpl
	FieldBlock *ObjectFieldBlock

	scope symbol.Scope
	obj   symbol.Symbol
}

var _ Scope = (*ObjectLiteral)(nil)
var _ Expr = (*ObjectLiteral)(nil)

func (s *ObjectLiteral) isExpr() {}

func (s *ObjectLiteral) Scope() symbol.Scope {
	return s.scope
}

func (s *ObjectLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		s.scope = symbol.NewScope()
		l := ctx.CurrentFunc().NewLocal()
		l.SetName("$obj")
		l.SetPos(s.Pos())
		if !s.scope.PutSymbol(ctx, l) {
			return
		}
		s.obj = l

	case Emit:
		// For completeness. Nothing can reference the $arr symbol, thus it cannot
		// be lifted, hence does not require initialization.
		emitVariableInit(ctx, s.Pos(), s.obj)
		emitSymbolRefPush(s.Pos(), ctx.Emitter(), s.obj)
		ctx.Emitter().Emit(s.Pos(), vm.OpNewObject, 0, 0)
		ctx.Emitter().Emit(s.Pos(), vm.OpStore, 1, 0)
	}

	ctx.RunPassChild(s, s.FieldBlock, pass)

	switch pass {
	case Emit:
		emitSymbolPush(s.Pos(), ctx.Emitter(), s.obj)
	}
}
