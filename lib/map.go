package lib

import (
	"context"

	"github.com/dcaiafa/nitro"
)

func fnMap(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	iter, err := getEnumeratorArg(ctx, args, 0)
	if err != nil {
		return nil, err
	}
	fn, err := getCallableArg(ctx, args, 1)
	if err != nil {
		return nil, err
	}

	mapIter := &mapIter{
		iter: iter,
		fn:   fn,
	}

	c := nitro.NewClosure(mapIter.Next, nil)
	return []nitro.Value{c}, nil
}

type mapIter struct {
	iter nitro.Value
	fn   nitro.Value
}

func (i *mapIter) Next(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	v, ok, err := nitro.Next(ctx, i.iter, 1)
	if err != nil {
		return nil, err
	}
	if !ok {
		return []nitro.Value{nitro.NewBool(false), nil}, nil
	}
	res, err := nitro.Call(ctx, i.fn, v, 1)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewBool(true), res[0]}, nil
}
