package ast

import (
	"fmt"
	"reflect"

	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/token"
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

	emitter *runtime.Emitter
	errors  []error
}

func NewContext() *Context {
	return &Context{
		emitter: runtime.NewEmitter(),
	}
}

func (s *Context) Get(outAST *AST) bool {
	outType := reflect.TypeOf(*outAST)
	for i := len(s.stack) - 1; i >= 0; i-- {
		ast := s.stack[i]
		if reflect.TypeOf(ast).AssignableTo(outType) {
			*outAST = ast
			return true
		}
	}
	return false
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

func (s *Context) CurrentFunc() *types.FuncSymbol {
	for i := len(s.stack) - 1; i >= 0; i-- {
		ast := s.stack[i]
		if funcAST, ok := ast.(Func); ok {
			return funcAST.Func()
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

func (c *Context) HasErrors() bool {
	return len(c.errors) != 0
}

func (c *Context) Failf(pos token.Pos, msg string, args ...interface{}) {
	c.errors = append(c.errors, fmt.Errorf(msg, args...))
}

func (c *Context) Errors() []error {
	return c.errors
}

type PassRunner interface {
	RunPass(ctx *Context, pass Pass)
}
