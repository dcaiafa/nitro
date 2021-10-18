package nitroco

import (
	"context"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
)

var errRunWithTimeoutUsage = nitro.NewInvalidUsageError("run_with_timeout(callable, dur)")

func runWithTimeout(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errRunWithTimeoutUsage
	}

	dur, ok := args[1].(lib.Duration)
	if !ok {
		return nil, errRunWithTimeoutUsage
	}

	ctx, cancel := context.WithTimeout(vm.Context(), dur.Duration())
	defer cancel()

	vm.PushContext(ctx)
	defer vm.PopContext()

	_, err := vm.Call(args[0], nil, 0)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
