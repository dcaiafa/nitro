package ast

import "github.com/dcaiafa/nitro/internal/vm"

type DeferStmt struct {
	PosImpl

	deferred *deferred
}

func NewDeferStmt(callExpr *FuncCallExpr) *DeferStmt {
	return &DeferStmt{
		deferred: newDeferred(callExpr),
	}
}

func (s *DeferStmt) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.deferred, pass)

	if pass == Emit {
		ctx.Emitter().Emit(s.Pos(), vm.OpDefer, 0, 0)
	}
}

type deferred struct {
	Func
}

func newDeferred(callExpr *FuncCallExpr) *deferred {
	d := new(deferred)
	d.Block = &StmtBlock{
		Stmts: ASTs{
			&ExprStmt{Expr: callExpr},
		},
	}
	d.DebugName = "$defer"
	d.IsClosure = true
	return d
}

func (d *deferred) RunPass(ctx *Context, pass Pass) {
	d.Func.RunPass(ctx, pass)
}
