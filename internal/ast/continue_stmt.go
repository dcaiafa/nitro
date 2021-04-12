package ast

type ContinueStmt struct {
	astBase
}

func (s *ContinueStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Emit {
		var continuer BreakContinueEmitter
	Loop:
		for i := 0; i < ctx.Len(); i++ {
			switch ast := ctx.Peek(i).(type) {
			case BreakContinueEmitter:
				continuer = ast
			case *Func:
				break Loop
			}
		}
		if continuer == nil {
			ctx.Failf(
				s.Pos(), "'break' can only be used from inside a 'while' or 'for' loop")
			return
		}
		continuer.EmitContinue(s.Pos(), ctx.Emitter())
	}
}
