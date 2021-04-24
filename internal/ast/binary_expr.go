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

func binOpAstToRuntime(op Operator) vm.BinOp {
	switch op {
	case BinOpPlus:
		return vm.BinAdd
	case BinOpMinus:
		return vm.BinSub
	case BinOpMult:
		return vm.BinMult
	case BinOpDiv:
		return vm.BinDiv
	case BinOpMod:
		return vm.BinMod
	case BinOpLT:
		return vm.BinLT
	case BinOpLE:
		return vm.BinLE
	case BinOpGT:
		return vm.BinGT
	case BinOpGE:
		return vm.BinGE
	case BinOpEq:
		return vm.BinEq
	case BinOpNE:
		return vm.BinNE
	default:
		panic("invalid operator")
	}
}
