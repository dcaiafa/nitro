package context

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/typecheck"
)

type Context struct {
	scopes *typecheck.ScopeStack
	errors []error
}

func NewContext() *Context {
	return &Context{
		scopes: typecheck.NewScopeStack(),
	}
}

func (c *Context) Scopes() *typecheck.ScopeStack {
	return c.scopes
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
