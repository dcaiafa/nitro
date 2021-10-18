package nitroco

import "github.com/dcaiafa/nitro"

func RegisterNativePackage(c *nitro.Compiler) {
	c.RegisterNativeModuleLoader("co", func(r nitro.NativeModuleContext) {
		r.RegisterNativeFn("start", start)
		r.RegisterNativeFn("run_with_timeout", runWithTimeout)
	})
}
