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
}

type ASTs []AST

func (asts ASTs) RunPass(ctx *context.Context, pass context.Pass) {
	for _, ast := range asts {
		ast.RunPass(ctx, pass)
	}
}

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

type ExprStmt struct {
	Expr Expr
}

func (s *ExprStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type WhileStmt struct {
	Predicate Expr
	Stmts     ASTs
}

func (s *WhileStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type FuncParam struct {
	Name token.Token
}

func (s *FuncParam) RunPass(ctx *context.Context, pass context.Pass) {
}

type IfStmt struct {
	Blocks ASTs
}

func (s *IfStmt) RunPass(ctx *context.Context, pass context.Pass) {
	for _, block := range s.Blocks {
		block.RunPass(ctx, pass)
	}
}

type IfBlock struct {
	Pred  Expr
	Stmts ASTs
}

func (s *IfBlock) RunPass(ctx *context.Context, pass context.Pass) {
}

type ForStmt struct {
	ForVars  ASTs
	IterExpr Expr
	Stmts    ASTs
}

func (s *ForStmt) RunPass(ctx *context.Context, pass context.Pass) {
}

type ForVar struct {
	VarName token.Token
}

func (s *ForVar) RunPass(ctx *context.Context, pass context.Pass) {
}

type AssignStmt struct {
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
	Term Expr
	Op   Operator
}

func (s *UnaryExpr) Value() *Value {
	return nil
}

func (s *UnaryExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type LiteralExpr struct {
	Val token.Token
}

func (s *LiteralExpr) Value() *Value {
	return nil
}

func (s *LiteralExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type IndexExpr struct {
	Target Expr
	Index  Expr
}

func (s *IndexExpr) Value() *Value {
	return nil
}

func (s *IndexExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type MemberAccess struct {
	Target Expr
	Member token.Token
}

func (s *MemberAccess) Value() *Value {
	return nil
}

func (s *MemberAccess) RunPass(ctx *context.Context, pass context.Pass) {
}

type FuncCall struct {
	Target Expr
	Args   Exprs
}

func (s *FuncCall) Value() *Value {
	return nil
}

func (s *FuncCall) RunPass(ctx *context.Context, pass context.Pass) {
}

type LambdaExpr struct {
	FuncParams ASTs
	Stmts      ASTs
}

func (s *LambdaExpr) Value() *Value {
	return nil
}

func (s *LambdaExpr) RunPass(ctx *context.Context, pass context.Pass) {
}

type ObjectLiteral struct {
	Fields ASTs
}

func (s *ObjectLiteral) Value() *Value {
	return nil
}

func (s *ObjectLiteral) RunPass(ctx *context.Context, pass context.Pass) {
}

type ObjectField struct {
	NameID   token.Token
	NameExpr Expr
	Val      Expr
}

func (s *ObjectField) RunPass(ctx *context.Context, pass context.Pass) {
}

type ArrayLiteral struct {
	Elements ASTs
}

func (s *ArrayLiteral) Value() *Value {
	return nil
}

func (s *ArrayLiteral) RunPass(ctx *context.Context, pass context.Pass) {
}

type ArrayElement struct {
	Val Expr
}

func (s *ArrayElement) RunPass(ctx *context.Context, pass context.Pass) {
}

type ReturnStmt struct {
	Values Exprs
}

func (s *ReturnStmt) RunPass(ctx *context.Context, pass context.Pass) {
}
