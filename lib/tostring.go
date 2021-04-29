package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errToStringUsage = errors.New(
	`invalid usage. Expected tostring(any)`)

func tostring(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errToStringUsage
	}

	if args[0] == nil {
		return []nitro.Value{nitro.NewString("<nil>")}, nil
	}

	return []nitro.Value{nitro.NewString(args[0].String())}, nil
}
