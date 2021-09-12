package lib

import (
	"encoding/hex"
	"errors"

	"github.com/dcaiafa/nitro"
)

var errToHexUsage = errors.New(
	`invalid usage. Expected tohex(string)`)

func tohex(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errToHexUsage
	}

	input, ok := args[0].(nitro.String)
	if !ok {
		return nil, errSha1Usage
	}

	res := hex.EncodeToString([]byte(input.String()))
	return []nitro.Value{nitro.NewString(res)}, nil
}
