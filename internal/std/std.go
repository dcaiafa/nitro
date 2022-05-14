package std

import "github.com/dcaiafa/nitro/internal/vm"

type Compiler interface {
	AddNativeFn(name string, fn func(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error))
}

func Register(c Compiler) {
	c.AddNativeFn("args", args)
	c.AddNativeFn("delete", delete_)
	c.AddNativeFn("has", has)
	c.AddNativeFn("len", Len)
	c.AddNativeFn("narg", narg)
	c.AddNativeFn("range", Range)
}
