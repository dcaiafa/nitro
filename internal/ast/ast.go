package ast

import (
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/token"
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
	Scope() *symbol.Scope
}

type astBase struct {
	pos token.Pos
}

func (b *astBase) Pos() token.Pos {
	return b.pos
}

func (b *astBase) SetPos(pos token.Pos) {
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

type WhileStmt struct {
	astBase
	Predicate Expr
	Stmts     ASTs
}

func (s *WhileStmt) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}
