package std

import "github.com/dcaiafa/nitro/internal/runtime"

type Compiler interface {
	AddExternalFn(name string, fn runtime.ExternFn)
}

func Register(c Compiler) {
	c.AddExternalFn("narg", narg)
	c.AddExternalFn("len", Len)
	c.AddExternalFn("range", Range)
	c.AddExternalFn("push", push)
	c.AddExternalFn("has", has)
	c.AddExternalFn("delete", delete_)
}
