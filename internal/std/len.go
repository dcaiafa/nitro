package std

import (
	"context"
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

var errNotEnoughArgs = errors.New("not enough arguments")

func Len(ctx context.Context, args []runtime.Value) ([]runtime.Value, error) {
	if len(args) == 0 {
		return nil, errNotEnoughArgs
	}

	v := args[0]
	if v == nil {
		return []runtime.Value{runtime.Int(0)}, nil
	}

	var l int
	switch v := v.(type) {
	case runtime.String:
		l = len(v)
	case *runtime.Array:
		l = v.Len()
	case *runtime.Object:
		l = v.Len()
	default:
		return nil, fmt.Errorf("cannot get length of %s", v.ValueType())
	}

	return []runtime.Value{runtime.Int(l)}, nil
}
