package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/types"
)

type Pass int

const (
	Print Pass = iota

	CreateAndResolveNames
	Emit
)

type Context struct {
	Stack
	*errlogger.ErrLoggerBase

	emitter *runtime.Emitter
}

func NewContext(l errlogger.ErrLogger) *Context {
	return &Context{
		ErrLoggerBase: errlogger.NewErrLoggerBase(l),
		emitter:       runtime.NewEmitter(),
	}
}

func (c *Context) RunPassChild(parent AST, child AST, pass Pass) {
	c.Push(parent)
	child.RunPass(c, pass)
	c.Pop()
}

func (s *Context) Parent() AST {
	return s.stack[len(s.stack)-1]
}

func (s *Context) FindSymbol(symName string) types.Symbol {
	for i := len(s.stack) - 1; i >= 0; i-- {
		ast := s.stack[i]
		if scope, ok := ast.(Scope); ok {
			sym := scope.Scope().GetSymbol(symName)
			if sym != nil {
				return sym
			}
		}
	}
	return nil
}

func (s *Context) CurrentFunc() *Func {
	for i := len(s.stack) - 1; i >= 0; i-- {
		ast := s.stack[i]
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

func (s *Context) CurrentScope() *types.Scope {
	for i := len(s.stack) - 1; i >= 0; i-- {
		ast := s.stack[i]
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
