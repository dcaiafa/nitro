package std

import "github.com/dcaiafa/nitro/internal/runtime"

type Compiler interface {
	AddExternalFn(name string, fn runtime.ExternFn)
}

func Register(c Compiler) {
	c.AddExternalFn("len", Len)
	c.AddExternalFn("range", Range)
	c.AddExternalFn("push", fnPush)
}
