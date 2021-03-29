package lib

import "github.com/dcaiafa/nitro"

func RegisterAll(c *nitro.Compiler) {
	c.AddExternalFn("in", In)
}
