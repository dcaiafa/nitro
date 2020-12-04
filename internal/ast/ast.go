package ast

import (
	"github.com/dcaiafa/nitro/internal/context"
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
	context.PassRunner
	Pos() token.Pos
	SetPos(pos token.Pos)
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

func (asts ASTs) RunPass(ctx *context.Context, pass context.Pass) {
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
	Value() *Value
}

type Exprs []Expr

func (exprs Exprs) RunPass(ctx *context.Context, pass context.Pass) {
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

type ExprStmt struct {
	astBase
	Expr Expr
}

func (s *ExprStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type WhileStmt struct {
	astBase
	Predicate Expr
	Stmts     ASTs
}

func (s *WhileStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type FuncParam struct {
	astBase
	Name token.Token
}

func (s *FuncParam) RunPass(ctx *context.Context, pass context.Pass) {
}

type IfStmt struct {
	astBase
	Blocks ASTs
}

func (s *IfStmt) RunPass(ctx *context.Context, pass context.Pass) {
	for _, block := range s.Blocks {
		block.RunPass(ctx, pass)
	}
}

type IfBlock struct {
	astBase
	Pred  Expr
	Stmts ASTs
}

func (s *IfBlock) RunPass(ctx *context.Context, pass context.Pass) {
}

type ForStmt struct {
	astBase
	ForVars  ASTs
	IterExpr Expr
	Stmts    ASTs
}

func (s *ForStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type ForVar struct {
	astBase
	VarName token.Token
}

func (s *ForVar) RunPass(ctx *context.Context, pass context.Pass) {
}

type AssignStmt struct {
	astBase
	Lvalue Expr
	Rvalue Expr
}

func (s *AssignStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type Operator int

const (
	OpPlus Operator = iota
	OpMinus
	OpMult
	OpDiv
	OpLT
	OpLE
	OpGT
	OpGE
	OpEq
	OpNE
	OpAnd
	OpOr
	OpNot
)

type BinaryExpr struct {
	astBase
	Left  Expr
	Op    Operator
	Right Expr
}

func (s *BinaryExpr) Value() *Value {
	return nil
}

func (s *BinaryExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type UnaryExpr struct {
	astBase
	Term Expr
	Op   Operator
}

func (s *UnaryExpr) Value() *Value {
	return nil
}

func (s *UnaryExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type LiteralExpr struct {
	astBase
	Val token.Token
}

func (s *LiteralExpr) Value() *Value {
	return nil
}

func (s *LiteralExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type IndexExpr struct {
	astBase
	Target Expr
	Index  Expr
}

func (s *IndexExpr) Value() *Value {
	return nil
}

func (s *IndexExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type MemberAccess struct {
	astBase
	Target Expr
	Member token.Token
}

func (s *MemberAccess) Value() *Value {
	return nil
}

func (s *MemberAccess) RunPass(ctx *context.Context, pass context.Pass) {
}

type FuncCall struct {
	astBase
	Target Expr
	Args   Exprs
}

func (s *FuncCall) Value() *Value {
	return nil
}

func (s *FuncCall) RunPass(ctx *context.Context, pass context.Pass) {
}

type LambdaExpr struct {
	astBase
	FuncParams ASTs
	Stmts      ASTs
}

func (s *LambdaExpr) Value() *Value {
	return nil
}

func (s *LambdaExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type ObjectLiteral struct {
	astBase
	Fields ASTs
}

func (s *ObjectLiteral) Value() *Value {
	return nil
}

func (s *ObjectLiteral) RunPass(ctx *context.Context, pass context.Pass) {
}

type ObjectField struct {
	astBase
	NameID   token.Token
	NameExpr Expr
	Val      Expr
}

func (s *ObjectField) RunPass(ctx *context.Context, pass context.Pass) {
}

type ArrayLiteral struct {
	astBase
	Elements ASTs
}

func (s *ArrayLiteral) Value() *Value {
	return nil
}

func (s *ArrayLiteral) RunPass(ctx *context.Context, pass context.Pass) {
}

type ArrayElement struct {
	astBase
	Val Expr
}

func (s *ArrayElement) RunPass(ctx *context.Context, pass context.Pass) {
}

type ReturnStmt struct {
	astBase
	Values Exprs
}

func (s *ReturnStmt) RunPass(ctx *context.Context, pass context.Pass) {
}
