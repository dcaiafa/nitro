package ast

type YieldStmt struct {
	astBase
	Values Exprs
}

func (s *YieldStmt) RunPass(ctx *Context, pass Pass) {
	if pass == Rewrite {
		fn := ctx.CurrentFunc()
		if fn == nil {
			ctx.Failf(s.Pos(), "cannot yield outside of a function")
			return
		}
		fn.MarkEnumerable()
	}
	ctx.RunPassChild(s, s.Values, pass)
}
