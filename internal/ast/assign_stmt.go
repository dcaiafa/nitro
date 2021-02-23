package ast

import "github.com/dcaiafa/nitro/internal/runtime"

type AssignStmt struct {
	astBase
	Lvalues ASTs
	Rvalue  Expr
}

func (s *AssignStmt) RunPass(ctx *Context, pass Pass) {
	ctx.RunPassChild(s, s.Lvalues, pass)
	ctx.RunPassChild(s, s.Rvalue, pass)

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.Emit(runtime.OpStore, uint16(len(s.Lvalues)), 0)
	}
}
