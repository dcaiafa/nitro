package lib

import (
	"os"

	"github.com/dcaiafa/nitro"
)

var errUserHomeDirUsage = nitro.NewInvalidUsageError("home_dir()")

func homeDir(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errUserHomeDirUsage
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewString(dir)}, nil
}
