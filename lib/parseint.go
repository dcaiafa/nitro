package lib

import (
	"errors"
	"strconv"

	"github.com/dcaiafa/nitro"
)

var errParseIntUsage = errors.New(
	`invalid usage. Expected parseint(str, map?)`)

type parseIntOptions struct {
	Base int64
}

var parseIntConv Value2Structer

func parseint(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errParseIntUsage
	}
	str, ok := args[0].(nitro.String)
	if !ok {
		return nil, errParseIntUsage
	}

	var opts parseIntOptions
	if len(args) == 2 {
		err := parseIntConv.Convert(args[1], &opts)
		if err != nil {
			return nil, err
		}
	}

	i, err := strconv.ParseInt(str.String(), int(opts.Base), 64)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(i)}, nil
}
