package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type ArrayLiteral struct {
	astBase
	Block *ArrayElementBlock

	scope *symbol.Scope
	arr   symbol.Symbol
}

func (a *ArrayLiteral) isExpr() {}

func (a *ArrayLiteral) Scope() *symbol.Scope {
	return a.scope
}

func (a *ArrayLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		a.scope = symbol.NewScope()
		ctx.Push(a)
		a.arr = AddVariable(ctx, "$arr", a.Pos())
		ctx.Pop()

	case Emit:
		emitSymbolRefPush(a.Pos(), ctx.Emitter(), a.arr)
		ctx.Emitter().Emit(a.Pos(), runtime.OpNewArray, 0, 0)
		ctx.Emitter().Emit(a.Pos(), runtime.OpStore, 1, 0)
	}

	ctx.RunPassChild(a, a.Block, pass)

	switch pass {
	case Emit:
		emitSymbolPush(a.Pos(), ctx.Emitter(), a.arr)
	}
}

type ArrayElementBlock struct {
	astBase

	Elements ASTs

	scope *symbol.Scope
}

func (b *ArrayElementBlock) Scope() *symbol.Scope {
	return b.scope
}

func (b *ArrayElementBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = symbol.NewScope()

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
		ctx.Emitter().Emit(b.Pos(), runtime.OpPop, 1, 0)
	}
}
