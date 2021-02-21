package ast

import (
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
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
	Scope() *types.Scope
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

type ForStmt struct {
	astBase
	ForVars  ASTs
	IterExpr Expr
	Stmts    ASTs
}

func (s *ForStmt) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type ForVar struct {
	astBase
	VarName token.Token
}

func (s *ForVar) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type UnaryOp int

const (
	UnaryOpNot UnaryOp = iota
	UnaryOpPlus
	UnaryOpMinus
)

type UnaryExpr struct {
	astBase
	Term Expr
	Op   UnaryOp
}

func (s *UnaryExpr) isExpr() {}

func (s *UnaryExpr) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type IndexExpr struct {
	astBase
	Target Expr
	Index  Expr
}

func (s *IndexExpr) isExpr() {}

func (s *IndexExpr) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type MemberAccess struct {
	astBase
	Target Expr
	Member token.Token
}

func (s *MemberAccess) isExpr() {}

func (s *MemberAccess) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type ArrayLiteral struct {
	astBase
	Elements ASTs
}

func (s *ArrayLiteral) isExpr() {}

func (s *ArrayLiteral) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}

type ArrayElement struct {
	astBase
	Val Expr
}

func (s *ArrayElement) RunPass(ctx *Context, pass Pass) {
	panic("not implemented")
}
