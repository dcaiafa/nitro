package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
	errTakesNoArgs   = errors.New("function does not take arguments")
)

func RegisterAll(c *nitro.Compiler) {
	c.AddNativeFn("parsecsv", parsecsv)
	c.AddNativeFn("avg", avg)
	c.AddNativeFn("buf", buf)
	c.AddNativeFn("close", closep)
	c.AddNativeFn("count", count)
	c.AddNativeFn("cp", cp)
	c.AddNativeFn("create", create)
	c.AddNativeFn("createtemp", createtemp)
	c.AddNativeFn("dur", dur)
	c.AddNativeFn("env", env)
	c.AddNativeFn("stderr", stderr)
	c.AddNativeFn("exec", exec)
	c.AddNativeFn("fileexists", fileexists)
	c.AddNativeFn("filter", filter)
	c.AddNativeFn("find", find)
	c.AddNativeFn("group", group)
	c.AddNativeFn("hasprefix", hasprefix)
	c.AddNativeFn("isdir", isdir)
	c.AddNativeFn("lines", lines)
	c.AddNativeFn("log", log)
	c.AddNativeFn("logf", logf)
	c.AddNativeFn("ls", ls)
	c.AddNativeFn("map", mapp)
	c.AddNativeFn("mapreduce", mapreduce)
	c.AddNativeFn("match", match)
	c.AddNativeFn("matchall", matchall)
	c.AddNativeFn("max", max)
	c.AddNativeFn("min", min)
	c.AddNativeFn("now", now)
	c.AddNativeFn("open", open)
	c.AddNativeFn("parsebase64", parsebase64)
	c.AddNativeFn("parseint", parseint)
	c.AddNativeFn("parsejson", parsejson)
	c.AddNativeFn("parsetime", parsetime)
	c.AddNativeFn("pathbase", pathbase)
	c.AddNativeFn("pathclean", pathclean)
	c.AddNativeFn("pathdir", pathdir)
	c.AddNativeFn("pathext", pathext)
	c.AddNativeFn("pathfromslash", pathfromslash)
	c.AddNativeFn("pathjoin", pathjoin)
	c.AddNativeFn("pathmatch", pathmatch)
	c.AddNativeFn("popstdout", popstdout)
	c.AddNativeFn("print", print)
	c.AddNativeFn("printall", printall)
	c.AddNativeFn("printf", printf)
	c.AddNativeFn("prompt", prompt)
	c.AddNativeFn("pushstdout", pushstdout)
	c.AddNativeFn("read", read)
	c.AddNativeFn("readfile", readfile)
	c.AddNativeFn("reduce", reduce)
	c.AddNativeFn("regex", regex)
	c.AddNativeFn("replace", replace)
	c.AddNativeFn("rm", rm)
	c.AddNativeFn("rmall", rmall)
	c.AddNativeFn("seek", seek)
	c.AddNativeFn("shuffle", shuffle)
	c.AddNativeFn("skip", skip)
	c.AddNativeFn("sleep", sleep)
	c.AddNativeFn("sort", sort)
	c.AddNativeFn("split", split)
	c.AddNativeFn("sprintf", sprintf)
	c.AddNativeFn("stdin", stdin)
	c.AddNativeFn("stdout", stdout)
	c.AddNativeFn("sum", sum)
	c.AddNativeFn("take", take)
	c.AddNativeFn("timefromunix", timefromunix)
	c.AddNativeFn("toarray", toarray)
	c.AddNativeFn("tobase64", tobase64)
	c.AddNativeFn("tojson", tojson)
	c.AddNativeFn("tolower", tolower)
	c.AddNativeFn("tomap", tomap)
	c.AddNativeFn("tostring", tostring)
	c.AddNativeFn("toupper", toupper)
	c.AddNativeFn("trim", trim)
	c.AddNativeFn("unique", unique)
	c.AddNativeFn("writefile", writefile)
	c.AddNativeFn("writeto", writeto)
}
