package ast

import "github.com/dcaiafa/nitro/internal/runtime"

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
		ctx.Emitter().Emit(e.Pos(), runtime.OpEvalBinOp, uint32(binOpAstToRuntime(e.Op)), 0)
	}
}

func binOpAstToRuntime(op Operator) runtime.BinOp {
	switch op {
	case BinOpPlus:
		return runtime.BinAdd
	case BinOpMinus:
		return runtime.BinSub
	case BinOpMult:
		return runtime.BinMult
	case BinOpDiv:
		return runtime.BinDiv
	case BinOpMod:
		return runtime.BinMod
	case BinOpLT:
		return runtime.BinLT
	case BinOpLE:
		return runtime.BinLE
	case BinOpGT:
		return runtime.BinGT
	case BinOpGE:
		return runtime.BinGE
	case BinOpEq:
		return runtime.BinEq
	case BinOpNE:
		return runtime.BinNE
	default:
		panic("invalid operator")
	}
}
