package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("fromjson", fromjson)
	c.AddExternalFn("in", in)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("map", fnMap)
	c.AddExternalFn("open", open)
	c.AddExternalFn("out", out)
	c.AddExternalFn("print", fnPrint)
	c.AddExternalFn("printall", fnPrintAll)
	c.AddExternalFn("readall", readall)
	c.AddExternalFn("tojson", fnToJSON)
}
