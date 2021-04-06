package lib

import (
	"context"

	"github.com/dcaiafa/nitro"
)

func fnUnique(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	e, err := getEnumeratorArg(ctx, args, 0)
	if err != nil {
		return nil, err
	}

	set := make(map[nitro.Value]bool)

	for {
		v, ok, err := nitro.Next(ctx, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		set[v[0]] = true
	}

	arr := nitro.NewArrayWithCapacity(len(set))
	for v := range set {
		arr.Push(v)
	}

	return []nitro.Value{arr}, nil
}
