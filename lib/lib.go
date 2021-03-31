package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("stdin", stdin)
	c.AddExternalFn("stdout", stdout)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("print", print)
}
