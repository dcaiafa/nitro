package lib

import (
	"strconv"

	"github.com/dcaiafa/nitro"
)

var errToIntUsage = nitro.NewInvalidUsageError("toint(string|float|int)")

func toint(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errToIntUsage
	}

	var res nitro.Int
	switch arg := args[0].(type) {
	case nitro.String:
		i, err := strconv.ParseInt(arg.String(), 0, 64)
		if err != nil {
			return nil, err
		}
		res = nitro.NewInt(i)

	case nitro.Float:
		res = nitro.NewInt(int64(arg.Float64()))

	case nitro.Int:
		res = arg

	default:
		return nil, errToIntUsage
	}

	return []nitro.Value{res}, nil
}
