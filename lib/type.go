package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errTypeUsage = errors.New(
	`invalid usage. Expected type(any)`)

func typep(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errTypeUsage
	}

	res := nitro.TypeName(args[0])

	return []nitro.Value{nitro.NewString(res)}, nil
}
