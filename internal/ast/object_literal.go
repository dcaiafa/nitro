package ast

import (
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type ObjectLiteral struct {
	astBase
	FieldBlock *ObjectFieldBlock

	scope *symbol.Scope
	obj   symbol.Symbol
}

func (s *ObjectLiteral) isExpr() {}

func (s *ObjectLiteral) Scope() *symbol.Scope {
	return s.scope
}

func (s *ObjectLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		s.scope = symbol.NewScope()
		ctx.Push(s)
		s.obj = AddVariable(ctx, "$obj", s.Pos())
		ctx.Pop()

	case Emit:
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
