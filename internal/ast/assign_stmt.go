package ast

import "github.com/dcaiafa/nitro/internal/vm"

type AssignStmt struct {
	PosImpl
	Lvalues ASTs
	Rvalues Exprs
}

func (s *AssignStmt) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if len(s.Lvalues) != len(s.Rvalues) {
			if funcCall, ok := s.Rvalues[0].(*FuncCallExpr); ok && len(s.Rvalues) == 1 {
				funcCall.RetN = len(s.Lvalues)
			} else {
				ctx.Failf(
					s.Pos(),
					"Left side of assignment expects %v values, "+
						"but right side produces %v values",
					len(s.Lvalues), len(s.Rvalues))
			}
		}
	}

	ctx.RunPassChild(s, s.Lvalues, pass)
	ctx.RunPassChild(s, s.Rvalues, pass)

	switch pass {
	case Emit:
		emitter := ctx.Emitter()
		emitter.Emit(s.Pos(), vm.OpStore, uint32(len(s.Lvalues)), 0)
	}
}
