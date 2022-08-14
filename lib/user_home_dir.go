package lib

import (
	"os"

	"github.com/dcaiafa/nitro"
)

func homeDir(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 0, 0); err != nil {
		return nil, err
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewString(dir)}, nil
}
