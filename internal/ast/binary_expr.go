package ast

import "github.com/dcaiafa/nitro/internal/vm"

type Operator int

const (
	BinOpPlus Operator = iota
	BinOpMinus
	BinOpMult
	BinOpDiv
	BinOpMod
	BinOpLT
	BinOpLE
	BinOpGT
	BinOpGE
	BinOpEq
	BinOpNE
)

type BinaryExpr struct {
	astBase
	Left  Expr
	Op    Operator
	Right Expr
}

func (e *BinaryExpr) isExpr() {}

func (e *BinaryExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Left, pass)
	ctx.RunPassChild(e, e.Right, pass)

	switch pass {
	case Emit:
		ctx.Emitter().Emit(e.Pos(), vm.OpEvalBinOp, uint32(binOpAstToRuntime(e.Op)), 0)
	}
}

func binOpAstToRuntime(op Operator) vm.Op {
	switch op {
	case BinOpPlus:
		return vm.OpAdd
	case BinOpMinus:
		return vm.OpSub
	case BinOpMult:
		return vm.OpMult
	case BinOpDiv:
		return vm.OpDiv
	case BinOpMod:
		return vm.OpMod
	case BinOpLT:
		return vm.OpLT
	case BinOpLE:
		return vm.OpLE
	case BinOpGT:
		return vm.OpGT
	case BinOpGE:
		return vm.OpGE
	case BinOpEq:
		return vm.OpEq
	case BinOpNE:
		return vm.OpNE
	default:
		panic("invalid operator")
	}
}
