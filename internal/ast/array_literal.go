package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type ArrayLiteral struct {
	PosImpl
	Block *ArrayElementBlock

	scope scope.Scope
	arr   symbol.Symbol
}

var _ Scope = (*ArrayLiteral)(nil)

func (a *ArrayLiteral) isExpr() {}

func (a *ArrayLiteral) Scope() scope.Scope {
	return a.scope
}

func (a *ArrayLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		a.scope = scope.NewScope(scope.Block)
		l := ctx.CurrentFunc().NewLocal()
		l.SetName("$arr")
		l.SetPos(a.Pos())
		if !a.scope.PutSymbol(ctx, l) {
			return
		}
		a.arr = l

	case Emit:
		// For completeness. Nothing can reference the $arr symbol, thus it cannot
		// be lifted, hence does not require initialization.
		emitVariableInit(ctx, a.Pos(), a.arr)
		emitSymbolRefPush(a.Pos(), ctx.Emitter(), a.arr)
		ctx.Emitter().Emit(a.Pos(), vm.OpNewArray, 0, 0)
		ctx.Emitter().Emit(a.Pos(), vm.OpStore, 1, 0)
	}

	ctx.RunPassChild(a, a.Block, pass)

	switch pass {
	case Emit:
		emitSymbolPush(a.Pos(), ctx.Emitter(), a.arr)
	}
}

type ArrayElementBlock struct {
	PosImpl

	Elements ASTs

	scope scope.Scope
}

var _ Scope = (*ArrayElementBlock)(nil)

func (b *ArrayElementBlock) Scope() scope.Scope {
	return b.scope
}

func (b *ArrayElementBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = scope.NewScope(scope.Block)

	case Emit:
		obj := ctx.FindSymbol("$arr")
		if obj == nil {
			panic("not reached")
		}
		emitSymbolPush(b.Pos(), ctx.Emitter(), obj)
	}

	ctx.Push(b)
	for _, b := range b.Elements {
		b.RunPass(ctx, pass)
	}
	ctx.Pop()

	switch pass {
	case Emit:
		ctx.Emitter().Emit(b.Pos(), vm.OpPop, 1, 0)
	}
}
