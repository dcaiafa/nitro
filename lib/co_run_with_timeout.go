package lib

import (
	"context"

	"github.com/dcaiafa/nitro"
)

func runWithTimeout(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	dur, err := getDurationArg(args, 1)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(vm.Context(), dur.Duration())
	defer cancel()

	vm.PushContext(ctx)
	defer vm.PopContext()

	_, err = vm.Call(args[0], nil, 0)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
