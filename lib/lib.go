package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
)

func RegisterAll(c *nitro.Compiler) {
	c.AddNativeFn("avg", avg)
	c.AddNativeFn("close", closep)
	c.AddNativeFn("cp", cp)
	c.AddNativeFn("create", create)
	c.AddNativeFn("createtemp", createtemp)
	c.AddNativeFn("exec", exec)
	c.AddNativeFn("fileexists", fileexists)
	c.AddNativeFn("filter", filter)
	c.AddNativeFn("fromjson", fromjson)
	c.AddNativeFn("hasprefix", hasprefix)
	c.AddNativeFn("in", in)
	c.AddNativeFn("isdir", isdir)
	c.AddNativeFn("lines", lines)
	c.AddNativeFn("log", log)
	c.AddNativeFn("logf", logf)
	c.AddNativeFn("ls", ls)
	c.AddNativeFn("map", mapp)
	c.AddNativeFn("mapreduce", mapreduce)
	c.AddNativeFn("match", match)
	c.AddNativeFn("max", max)
	c.AddNativeFn("open", open)
	c.AddNativeFn("out", out)
	c.AddNativeFn("parsetime", parsetime)
	c.AddNativeFn("pathbase", pathbase)
	c.AddNativeFn("pathclean", pathclean)
	c.AddNativeFn("pathdir", pathdir)
	c.AddNativeFn("pathext", pathext)
	c.AddNativeFn("pathfromslash", pathfromslash)
	c.AddNativeFn("pathjoin", pathjoin)
	c.AddNativeFn("popout", popout)
	c.AddNativeFn("print", print)
	c.AddNativeFn("printall", printall)
	c.AddNativeFn("printf", printf)
	c.AddNativeFn("pushout", pushout)
	c.AddNativeFn("read", read)
	c.AddNativeFn("replace", replace)
	c.AddNativeFn("rm", rm)
	c.AddNativeFn("rmall", rmall)
	c.AddNativeFn("shuffle", shuffle)
	c.AddNativeFn("skip", skip)
	c.AddNativeFn("sort", sort)
	c.AddNativeFn("split", split)
	c.AddNativeFn("sprintf", sprintf)
	c.AddNativeFn("take", take)
	c.AddNativeFn("timefromunix", timefromunix)
	c.AddNativeFn("timetounix", timetounix)
	c.AddNativeFn("timetounixnano", timetounixnano)
	c.AddNativeFn("toarray", toarray)
	c.AddNativeFn("tojson", tojson)
	c.AddNativeFn("top", top)
	c.AddNativeFn("trim", trim)
	c.AddNativeFn("unique", unique)
	c.AddNativeFn("write", write)
}
