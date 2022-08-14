package lib

import (
	"math"

	"github.com/dcaiafa/nitro"
)

func mathTrunc(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	v, err := getFloatArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := math.Trunc(v)
	return []nitro.Value{nitro.NewFloat(res)}, nil
}
