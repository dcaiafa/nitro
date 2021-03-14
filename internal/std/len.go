package std

import (
	"context"
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
)

func Len(ctx context.Context, caps []runtime.ValueRef, args []runtime.Value) ([]runtime.Value, error) {
	if len(args) == 0 {
		return nil, errNotEnoughArgs
	}

	v := args[0]
	if v == nil {
		return []runtime.Value{runtime.NewInt(0)}, nil
	}

	var l int
	switch v := v.(type) {
	case runtime.String:
		l = len(v.String())
	case *runtime.Array:
		l = v.Len()
	case *runtime.Object:
		l = v.Len()
	default:
		return nil, fmt.Errorf("cannot get length of %s", v.Type())
	}

	return []runtime.Value{runtime.NewInt(int64(l))}, nil
}
