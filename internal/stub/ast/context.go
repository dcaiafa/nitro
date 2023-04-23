package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/stub/analysis"
)

type Pass int

const (
	Check Pass = iota
	Emit
)

type Context struct {
	*errlogger.ErrLoggerWrapper

	Analysis *analysis.Analysis
}

func NewContext(l *errlogger.ErrLoggerWrapper) *Context {
	return &Context{
		ErrLoggerWrapper: l,
		Analysis:         analysis.NewAnalysis(),
	}
}

func (c *Context) RunPassChild(parent AST, child AST, pass Pass) {
	child.RunPass(c, pass)
}
