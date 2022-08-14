package lib

import "github.com/dcaiafa/nitro"

func toBool(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	res := nitro.CoerceToBool(args[0])

	return []nitro.Value{nitro.NewBool(res)}, nil
}
