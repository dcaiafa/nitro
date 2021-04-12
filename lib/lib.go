package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("filter", filter)
	c.AddExternalFn("fromjson", fromjson)
	c.AddExternalFn("in", in)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("map", fnMap)
	c.AddExternalFn("readall", readall)
	c.AddExternalFn("split", split)
	c.AddExternalFn("tojson", tojson)
	c.AddExternalFn("sort", sort)
	c.AddExternalFn("toarray", toarray)
	c.AddExternalFn("exec", exec)
	c.AddExternalFn("unique", unique)
	c.AddExternalFn("skip", skip)

	c.AddExternalFn("match", fnMatch)
	c.AddExternalFn("replace", replace)

	c.AddExternalFn("trim", trim)

	c.AddExternalFn("out", out)
	c.AddExternalFn("pushout", pushout)
	c.AddExternalFn("popout", popout)

	c.AddExternalFn("print", print)
	c.AddExternalFn("printf", printf)
	c.AddExternalFn("printall", printall)
	c.AddExternalFn("sprintf", sprintf)

	c.AddExternalFn("open", open)
	c.AddExternalFn("close", close_)

	c.AddExternalFn("fcreate", fcreate)
	c.AddExternalFn("fremove", fremove)
	c.AddExternalFn("fremoveall", fremoveall)
	c.AddExternalFn("fcreatetemp", fcreatetemp)
	c.AddExternalFn("fexists", fexists)
	c.AddExternalFn("flist", flist)
	c.AddExternalFn("fcopy", fcopy)
	c.AddExternalFn("fpathbase", fpathbase)
	c.AddExternalFn("fpathclean", fpathclean)
	c.AddExternalFn("fpathdir", fpathdir)
	c.AddExternalFn("fpathfromslash", fpathfromslash)
	c.AddExternalFn("fpathjoin", fpathjoin)
}
