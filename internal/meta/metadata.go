package meta

import (
	"github.com/dcaiafa/nitro/internal/runtime"
)

type Metadata struct {
	Params []Param
}

type Param struct {
	Name     string
	Type     Type
	Desc     string
	Default  runtime.Value
	Required bool
}
