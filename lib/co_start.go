package lib

import (
	"github.com/dcaiafa/nitro"
)

var errStartUsage error = nitro.NewInvalidUsageError("start(callable)")

func start(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errStartUsage
	}

	err := vm.StartCoroutine(args[0])
	if err != nil {
		return nil, err
	}

	return nil, nil
}
