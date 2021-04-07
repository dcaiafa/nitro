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
	c.AddExternalFn("close", fnClose)
	c.AddExternalFn("sort", fnSort)
	c.AddExternalFn("toarray", fnToArray)
	c.AddExternalFn("exec", fnExec)
	c.AddExternalFn("unique", fnUnique)
	c.AddExternalFn("skip", fnSkip)

	c.AddExternalFn("trim", fnTrim)

	c.AddExternalFn("out", fnOut)
	c.AddExternalFn("pushout", fnPushOut)
	c.AddExternalFn("popout", fnPopOut)

	c.AddExternalFn("print", fnPrint)
	c.AddExternalFn("printf", fnPrintf)
	c.AddExternalFn("printall", fnPrintAll)
	c.AddExternalFn("sprintf", fnSprintf)

	c.AddExternalFn("fopen", fnFOpen)
	c.AddExternalFn("fcreate", fnFCreate)
	c.AddExternalFn("fremove", fnFRemove)
	c.AddExternalFn("fremoveall", fnFRemoveAll)
	c.AddExternalFn("fcreatetemp", fnFCreateTemp)
	c.AddExternalFn("fexists", fnFExists)
	c.AddExternalFn("flist", fnFList)
	c.AddExternalFn("fcopy", fnFCopy)
	c.AddExternalFn("fpathbase", fnFPathBase)
	c.AddExternalFn("fpathclean", fnFPathClean)
	c.AddExternalFn("fpathdir", fnFPathDir)
	c.AddExternalFn("fpathfromslash", fnFPathFromSlash)
	c.AddExternalFn("fpathjoin", fnFPathJoin)
}
