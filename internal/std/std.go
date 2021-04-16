package std

import "github.com/dcaiafa/nitro/internal/runtime"

type Compiler interface {
	AddExternalFn(name string, fn runtime.ExternFn)
}

func Register(c Compiler) {
	c.AddExternalFn("args", args)
	c.AddExternalFn("delete", delete_)
	c.AddExternalFn("has", has)
	c.AddExternalFn("len", Len)
	c.AddExternalFn("narg", narg)
	c.AddExternalFn("push", push)
	c.AddExternalFn("range", Range)
}
