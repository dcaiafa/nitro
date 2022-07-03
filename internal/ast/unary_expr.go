package ast

import "github.com/dcaiafa/nitro/internal/vm"

type UnaryOp int

const (
	UnaryOpUndefined UnaryOp = iota
	UnaryOpNot
	UnaryOpPlus
	UnaryOpMinus
)

type UnaryExpr struct {
	PosImpl
	Term Expr
	Op   UnaryOp
}

func (e *UnaryExpr) isExpr() {}

func (e *UnaryExpr) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(e, e.Term, pass)

	switch pass {
	case Emit:
		switch e.Op {
		case UnaryOpNot:
			ctx.Emitter().Emit(e.Pos(), vm.OpNot, 0, 0)
		case UnaryOpPlus:
			// nop
		case UnaryOpMinus:
			ctx.Emitter().Emit(e.Pos(), vm.OpUnaryMinus, 0, 0)
		default:
			panic("not reached")
		}
	}
}
