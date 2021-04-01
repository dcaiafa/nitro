package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Pass int

const (
	Print Pass = iota

	Check
	Emit
)

type Context struct {
	Stack
	*errlogger.ErrLoggerWrapper

	emitter *runtime.Emitter
}

func NewContext(l *errlogger.ErrLoggerWrapper) *Context {
	return &Context{
		ErrLoggerWrapper: l,
		emitter:          runtime.NewEmitter(),
	}
}

func (c *Context) RunPassChild(parent AST, child AST, pass Pass) {
	if child == nil {
		return
	}

	c.Push(parent)
	child.RunPass(c, pass)
	c.Pop()
}

func (c *Context) Parent() AST {
	return c.stack[len(c.stack)-1]
}

func (c *Context) FindSymbol(symName string) symbol.Symbol {
	var fns []*Func
	var sym symbol.Symbol
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if scope, ok := ast.(Scope); ok {
			sym = scope.Scope().GetSymbol(symName)
			if sym != nil {
				break
			}
		}
		if fn, ok := ast.(*Func); ok {
			fns = append(fns, fn)
		}
	}
	if sym == nil {
		return nil
	}
	if len(fns) != 0 {
		if _, ok := sym.(symbol.Capturable); ok {
			for i := len(fns) - 1; i >= 0; i-- {
				sym = fns[i].NewCapture(sym)
			}
		}
	}
	return sym
}

func (c *Context) CurrentFunc() *Func {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if funcAST, ok := ast.(*Func); ok {
			return funcAST
		}
	}
	return nil
}

func (c *Context) Main() *Main {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if mainAST, ok := ast.(*Main); ok {
			return mainAST
		}
	}
	return nil
}

func (c *Context) CurrentScope() *symbol.Scope {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if scopeAST, ok := ast.(Scope); ok {
			return scopeAST.Scope()
		}
	}
	return nil
}

func (c *Context) Emitter() *runtime.Emitter {
	return c.emitter
}

type PassRunner interface {
	RunPass(ctx *Context, pass Pass)
}
