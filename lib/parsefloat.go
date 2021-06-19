package lib

import (
	"errors"
	"strconv"

	"github.com/dcaiafa/nitro"
)

var errParseFloatUsage = errors.New(
	`invalid usage. Expected parsefloat(str, map?)`)

func parsefloat(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errParseFloatUsage
	}
	str, ok := args[0].(nitro.String)
	if !ok {
		return nil, errParseFloatUsage
	}

	v, err := strconv.ParseFloat(str.String(), 64)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewFloat(v)}, nil
}
