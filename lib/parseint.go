package lib

import (
	"errors"
	"strconv"

	"github.com/dcaiafa/nitro"
)

var errParseIntUsage = errors.New(
	`invalid usage. Expected parseint(str)`)

func parseint(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errParseIntUsage
	}
	str, ok := args[0].(nitro.String)
	if !ok {
		return nil, errParseIntUsage
	}
	i, err := strconv.ParseInt(str.String(), 10, 64)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewInt(i)}, nil
}
