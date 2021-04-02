package lib

import (
	"errors"
	"io"

	"github.com/dcaiafa/nitro"
)

var errNotEnoughArgs = errors.New("not enough arguments")

type Reader struct {
	io.Reader
}

func (r *Reader) String() string { return "<Reader>" }
func (r *Reader) Type() string   { return "Reader" }

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("filter", fnFilter)
	c.AddExternalFn("fromjson", fromjson)
	c.AddExternalFn("in", in)
	c.AddExternalFn("lines", lines)
	c.AddExternalFn("map", fnMap)
	c.AddExternalFn("match", fnMatch)
	c.AddExternalFn("out", fnOut)
	c.AddExternalFn("print", fnPrint)
	c.AddExternalFn("printall", fnPrintAll)
	c.AddExternalFn("readall", readall)
	c.AddExternalFn("split", fnSplit)
	c.AddExternalFn("tojson", fnToJSON)
	c.AddExternalFn("pushout", fnPushOut)
	c.AddExternalFn("popout", fnPopOut)

	c.AddExternalFn("open", fnOpen)
	c.AddExternalFn("close", fnClose)
	c.AddExternalFn("create", fnCreate)
	c.AddExternalFn("remove", fnRemove)
	c.AddExternalFn("removeall", fnRemoveAll)
	c.AddExternalFn("createtemp", fnCreateTemp)
}
