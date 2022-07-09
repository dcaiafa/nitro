package lib

import (
	"math"

	"github.com/dcaiafa/nitro"
)

var errTruncUsage error = nitro.NewInvalidUsageError("trunc(float)")

func mathTrunc(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errTruncUsage
	}
	v, ok := args[0].(nitro.Float)
	if !ok {
		return nil, errTruncUsage
	}
	res := math.Trunc(v.Float64())
	return []nitro.Value{nitro.NewFloat(res)}, nil
}
