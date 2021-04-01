package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("in", in)
	c.AddExternalFn("out", out)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("print", print)
	c.AddExternalFn("fromjson", fromjson)
}
