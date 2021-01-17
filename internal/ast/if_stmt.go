package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type IfStmt struct {
	astBase
	Blocks ASTs

	end *runtime.Label
}

func (s *IfStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		s.end = ctx.Emitter().NewLabel()
	}

	ctx.RunPassChild(s, s.Blocks, pass)

	switch pass {
	case Emit:
		ctx.Emitter().AssignLabel(s.end)
	}
}

type IfBlock struct {
	astBase
	Pred  Expr
	Stmts ASTs

	end *runtime.Label
}

func (b *IfBlock) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		b.end = ctx.Emitter().NewLabel()
	}

	// The "else" block does not have a predicate expression.
	if b.Pred != nil {
		ctx.RunPassChild(b, b.Pred, pass)
	}

	switch pass {
	case Emit:
		if b.Pred != nil {
			ctx.Emitter().EmitJump(runtime.OpJumpIfFalse, b.end)
		}
	}

	ctx.RunPassChild(b, b.Stmts, pass)

	switch pass {
	case Emit:
		ifStmtEnd := ctx.Parent().(*IfStmt).end
		ctx.Emitter().EmitJump(runtime.OpJump, ifStmtEnd)
		ctx.Emitter().AssignLabel(b.end)
	}
}
