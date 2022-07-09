package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type ObjectFieldBlock struct {
	PosImpl

	Fields ASTs

	scope symbol.Scope
}

var _ Scope = (*ObjectFieldBlock)(nil)

func (b *ObjectFieldBlock) Scope() symbol.Scope {
	return b.scope
}

func (b *ObjectFieldBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		b.scope = symbol.NewScope()

	case Emit:
		obj := ctx.FindSymbol("$obj")
		if obj == nil {
			panic("not reached")
		}
		emitSymbolPush(b.Pos(), ctx.Emitter(), obj)
	}

	ctx.Push(b)
	for _, b := range b.Fields {
		b.RunPass(ctx, pass)
	}
	ctx.Pop()

	switch pass {
	case Emit:
		ctx.Emitter().Emit(b.Pos(), vm.OpPop, 1, 0)
	}
}

type ObjectField struct {
	PosImpl

	NameID   string
	NameExpr Expr
	Val      Expr
}

func (s *ObjectField) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.NameExpr, pass)

	switch pass {
	case Emit:
		if s.NameID != "" {
			emitter := ctx.Emitter()
			nameID := emitter.AddString(s.NameID)
			emitter.Emit(s.Pos(), vm.OpLoadLiteral, uint32(nameID), 0)
		}
	}

	ctx.RunPassChild(s, s.Val, pass)

	switch pass {
	case Emit:
		ctx.Emitter().Emit(s.Pos(), vm.OpObjectPutNoPop, 0, 0)
	}
}
