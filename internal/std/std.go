package std

import "github.com/dcaiafa/nitro/internal/runtime"

type Compiler interface {
	RegisterExternalFn(name string, fn runtime.ExternFn)
}

func Register(c Compiler) {
	c.RegisterExternalFn("len", Len)
}
