package meta

import (
	"github.com/dcaiafa/nitro/internal/runtime"
	"github.com/dcaiafa/nitro/internal/symbol"
)

type Metadata struct {
	Params []Param
}

type Param struct {
	Name     string
	Type     symbol.Type
	Desc     string
	Default  runtime.Value
	Required bool
}
