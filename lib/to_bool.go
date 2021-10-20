package lib

import "github.com/dcaiafa/nitro"

var errToBoolUsage = nitro.NewInvalidUsageError("to_bool(value)")

func toBool(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errToBoolUsage
	}

	res := nitro.CoerceToBool(args[0])

	return []nitro.Value{nitro.NewBool(res)}, nil
}
