package std

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func fnDelete(ctx context.Context, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return nil, nil
	}

	obj, ok := args[0].(interface {
		Delete(key runtime.Value)
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support 'delete'",
			runtime.TypeName(args[0]))
	}

	obj.Delete(args[1])
	return nil, nil
}
