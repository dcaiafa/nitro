package ast

import (
	"regexp"

	"github.com/dcaiafa/nitro/internal/vm"
)

type RegexLiteral struct {
	PosImpl

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
		idxLiteral := ctx.Emitter().AddLiteral(vm.NewRegex(r))
		ctx.Emitter().Emit(l.Pos(), vm.OpLoadLiteral, uint32(idxLiteral), 0)
	}
}
