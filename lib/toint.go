package lib

import (
	"strconv"

	"github.com/dcaiafa/nitro"
)

func toInt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
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
		return nil, errExpectedArg(0, args[0], "str", "float", "int")
	}

	return []nitro.Value{res}, nil
}
