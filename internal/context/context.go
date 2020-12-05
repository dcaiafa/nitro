package context

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/types"
)

type Context struct {
	scopes *types.ScopeStack
	errors []error
}

func NewContext() *Context {
	return &Context{
		scopes: types.NewScopeStack(),
	}
}

func (c *Context) Scopes() *types.ScopeStack {
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
