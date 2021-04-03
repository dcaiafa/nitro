package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("filter", fnFilter)
	c.AddExternalFn("fromjson", fromjson)
	c.AddExternalFn("in", in)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("map", fnMap)
	c.AddExternalFn("match", fnMatch)
	c.AddExternalFn("readall", readall)
	c.AddExternalFn("split", fnSplit)
	c.AddExternalFn("tojson", fnToJSON)

	c.AddExternalFn("out", fnOut)
	c.AddExternalFn("pushout", fnPushOut)
	c.AddExternalFn("popout", fnPopOut)

	c.AddExternalFn("print", fnPrint)
	c.AddExternalFn("printf", fnPrint)
	c.AddExternalFn("printall", fnPrintAll)

	c.AddExternalFn("fopen", fnFOpen)
	c.AddExternalFn("fclose", fnFClose)
	c.AddExternalFn("fcreate", fnFCreate)
	c.AddExternalFn("fremove", fnFRemove)
	c.AddExternalFn("fremoveall", fnFRemoveAll)
	c.AddExternalFn("fcreatetemp", fnFCreateTemp)
}
