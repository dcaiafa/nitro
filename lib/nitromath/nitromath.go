package nitromath

import "github.com/dcaiafa/nitro"

func RegisterNativePackage(c *nitro.Compiler) {
	c.RegisterNativeModuleLoader("math", func(r nitro.NativeModuleContext) {
		r.RegisterNativeFn("trunc", trunc)
	})
}
