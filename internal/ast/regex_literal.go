package ast

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"regexp"
)

type RegexLiteral struct {
	astBase

	RegexDef string
}

func (l *RegexLiteral) isExpr() {}

func (l *RegexLiteral) RunPass(ctx *Context, pass Pass) {
	if pass == Emit {
		regexDef := l.RegexDef[2 : len(l.RegexDef)-1]
		r, err := regexp.Compile(regexDef)
		if err != nil {
			ctx.Failf(l.Pos(), "Invalid regular expression: %v", err)
			return
		}
		idxLiteral := ctx.Emitter().AddLiteral(runtime.NewRegex(r))
		ctx.Emitter().Emit(l.Pos(), runtime.OpLoadLiteral, uint32(idxLiteral), 0)
	}
}
