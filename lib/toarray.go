package lib

import (
	"context"

	"github.com/dcaiafa/nitro"
)

func fnToArray(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	arr, err := ToArray(ctx, args[0])
	if err != nil {
		return nil, err
	}

	return []nitro.Value{arr}, nil
}

func ToArray(ctx context.Context, v nitro.Value) (*nitro.Array, error) {
	if arr, ok := v.(*nitro.Array); ok {
		return arr, nil
	}

	e, err := nitro.MakeEnumerator(ctx, v)
	if err != nil {
		return nil, err
	}

	arr := nitro.NewArray()
	for {
		v, ok, err := nitro.Next(ctx, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		arr.Push(v[0])
	}

	return arr, nil
}
