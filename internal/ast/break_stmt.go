package ast

type BreakStmt struct {
	astBase
}

func (s *BreakStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Emit {
		var breaker BreakContinueEmitter
	Loop:
		for i := 0; i < ctx.Len(); i++ {
			switch ast := ctx.Peek(i).(type) {
			case BreakContinueEmitter:
				breaker = ast
			case *Func:
				break Loop
			}
		}
		if breaker == nil {
			ctx.Failf(
				s.Pos(), "'break' can only be used from inside a 'while' or 'for' loop")
			return
		}
		breaker.EmitBreak(s.Pos(), ctx.Emitter())
	}
}
