package context

import (
	"github.com/dcaiafa/go-expr/expr/internal/symbol"
	"github.com/dcaiafa/go-expr/expr/runtime"
)

type Context struct {
	GlobalScope  *symbol.Scope
	Builder      *runtime.Builder
	GraphPrinter *GraphPrinter
}

func NewContext() *Context {
	return &Context{
		GlobalScope: symbol.NewScope(),
		Builder:     runtime.NewBuilder(),
	}
}
