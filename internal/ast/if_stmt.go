package ast

import "github.com/dcaiafa/nitro/internal/vm"

type IfStmt struct {
	astBase
	Sections ASTs

	end *vm.Label
}

func (s *IfStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Emit:
		s.end = ctx.Emitter().NewLabel()
	}

	ctx.RunPassChild(s, s.Sections, pass)

	switch pass {
	case Emit:
		ctx.Emitter().ResolveLabel(s.end)
	}
}

type IfBlock struct {
	astBase
	Pred  Expr
	Block AST

	end *vm.Label
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
			ctx.Emitter().EmitJump(b.Pos(), vm.OpJumpIfFalse, b.end, 0)
		}
	}

	ctx.RunPassChild(b, b.Block, pass)

	switch pass {
	case Emit:
		ifStmtEnd := ctx.Parent().(*IfStmt).end
		ctx.Emitter().EmitJump(b.Pos(), vm.OpJump, ifStmtEnd, 0)
		ctx.Emitter().ResolveLabel(b.end)
	}
}
