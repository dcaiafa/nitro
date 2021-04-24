package std

import "github.com/dcaiafa/nitro/internal/vm"

type Compiler interface {
	AddNativeFn(name string, fn vm.NativeFn)
}

func Register(c Compiler) {
	c.AddNativeFn("args", args)
	c.AddNativeFn("delete", delete_)
	c.AddNativeFn("has", has)
	c.AddNativeFn("len", Len)
	c.AddNativeFn("narg", narg)
	c.AddNativeFn("push", push)
	c.AddNativeFn("range", Range)
}
