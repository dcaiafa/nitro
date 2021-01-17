package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type Operator int

const (
	OpPlus Operator = iota
	OpMinus
	OpMult
	OpDiv
	OpMod
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

func (a *BinaryExpr) isExpr() {}

func (a *BinaryExpr) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		a.emit(ctx)

	default:
		a.Left.RunPass(ctx, pass)
		a.Right.RunPass(ctx, pass)
	}
}

func (a *BinaryExpr) emit(ctx *Context) {
	emitter := ctx.Emitter()
	switch a.Op {
	case OpPlus, OpMinus, OpMult, OpDiv, OpMod, OpLT, OpLE, OpGT, OpGE, OpEq, OpNE:
		a.Left.RunPass(ctx, Emit)
		a.Right.RunPass(ctx, Emit)
		emitter.Emit(runtime.OpBinOp, uint16(operatorToRuntime(a.Op)), 0)

	case OpAnd, OpOr:

	default:
		panic("invalid op")
	}
}

func operatorToRuntime(op Operator) runtime.BinOp {
	switch op {
	case OpPlus:
		return runtime.BinAdd
	case OpMinus:
		return runtime.BinSub
	case OpMult:
		return runtime.BinMult
	case OpDiv:
		return runtime.BinDiv
	case OpMod:
		return runtime.BinMod
	case OpLT:
		return runtime.BinLT
	case OpLE:
		return runtime.BinLE
	case OpGT:
		return runtime.BinGT
	case OpGE:
		return runtime.BinGE
	case OpEq:
		return runtime.BinEq
	case OpNE:
		return runtime.BinNE
	//case OpAnd:
	//case OpOr:
	default:
		panic("invalid operator")
	}
}
