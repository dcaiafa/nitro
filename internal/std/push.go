package std

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func fnPush(ctx context.Context, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	} else if len(args) > 2 {
		return nil, errTooManyArgs
	}

	arr, ok := args[0].(*runtime.Array)
	if !ok {
		return nil, fmt.Errorf(
			"expected Array in argument #1; instead got %q",
			runtime.TypeName(args[0]))
	}

	arr.Push(args[1])
	return nil, nil
}
