package lib

import (
	"errors"
	"os"

	"github.com/dcaiafa/nitro"
)

var errEnvUsage = errors.New(`invalid usage. Expected env(string?)`)

func env(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		allEnv := os.Environ()
		res := make([]nitro.Value, len(allEnv))
		for i, e := range allEnv {
			res[i] = nitro.NewString(e)
		}
		return []nitro.Value{nitro.NewArrayFromSlice(res)}, nil
	} else if len(args) == 1 {
		varName, ok := args[0].(nitro.String)
		if !ok {
			return nil, errEnvUsage
		}
		varValue := os.Getenv(varName.String())
		return []nitro.Value{nitro.NewString(varValue)}, nil
	} else {
		return nil, errEnvUsage
	}
}
