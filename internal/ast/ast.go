package ast

import (
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Type int

const (
	Number Type = iota
	String
	Bool
)

type Value struct {
	Type   Type
	Number float64
	String string
	Bool   bool
}

type AST interface {
	PassRunner
	Pos() token.Pos
	SetPos(pos token.Pos)
}

type Scope interface {
	Scope() scope.Scope
}

type RepeatableScope interface {
	Scope
	IsRepeatableScope()
}

type PosImpl struct {
	pos token.Pos
}

func (b *PosImpl) Pos() token.Pos {
	return b.pos
}

func (b *PosImpl) SetPos(pos token.Pos) {
	b.pos = pos
}

type ASTs []AST

func (asts ASTs) RunPass(ctx *Context, pass Pass) {
	for _, ast := range asts {
		ast.RunPass(ctx, pass)
	}
}

func (asts ASTs) Pos() token.Pos {
	if len(asts) > 0 {
		return asts[0].Pos()
	}
	return token.Pos{}
}

func (asts ASTs) SetPos(pos token.Pos) {}

type Expr interface {
	AST
	isExpr()
}

type Exprs []Expr

func (exprs Exprs) RunPass(ctx *Context, pass Pass) {
	for _, expr := range exprs {
		expr.RunPass(ctx, pass)
	}
}

func (exprs Exprs) Pos() token.Pos {
	if len(exprs) > 0 {
		return exprs[0].Pos()
	}
	return token.Pos{}
}

func (exprs Exprs) SetPos(pos token.Pos) {}

type BreakContinueEmitter interface {
	EmitBreak(pos token.Pos, emitter *vm.Emitter)
	EmitContinue(pos token.Pos, emitter *vm.Emitter)
}
