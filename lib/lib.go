package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("avg", avg)
	c.AddExternalFn("close", close_)
	c.AddExternalFn("exec", exec)
	c.AddExternalFn("fcopy", fcopy)
	c.AddExternalFn("fcreate", fcreate)
	c.AddExternalFn("fcreatetemp", fcreatetemp)
	c.AddExternalFn("fexists", fexists)
	c.AddExternalFn("filter", filter)
	c.AddExternalFn("flist", flist)
	c.AddExternalFn("fpathbase", fpathbase)
	c.AddExternalFn("fpathclean", fpathclean)
	c.AddExternalFn("fpathdir", fpathdir)
	c.AddExternalFn("fpathfromslash", fpathfromslash)
	c.AddExternalFn("fpathjoin", fpathjoin)
	c.AddExternalFn("fremove", fremove)
	c.AddExternalFn("fremoveall", fremoveall)
	c.AddExternalFn("fromjson", fromjson)
	c.AddExternalFn("in", in)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("map", fnMap)
	c.AddExternalFn("mapreduce", mapreduce)
	c.AddExternalFn("match", fnMatch)
	c.AddExternalFn("open", open)
	c.AddExternalFn("out", out)
	c.AddExternalFn("popout", popout)
	c.AddExternalFn("print", print)
	c.AddExternalFn("printall", printall)
	c.AddExternalFn("printf", printf)
	c.AddExternalFn("pushout", pushout)
	c.AddExternalFn("readall", readall)
	c.AddExternalFn("replace", replace)
	c.AddExternalFn("skip", skip)
	c.AddExternalFn("sort", sort)
	c.AddExternalFn("split", split)
	c.AddExternalFn("sprintf", sprintf)
	c.AddExternalFn("toarray", toarray)
	c.AddExternalFn("tojson", tojson)
	c.AddExternalFn("trim", trim)
	c.AddExternalFn("unique", unique)
}
