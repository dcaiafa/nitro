package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var errToMapUsage = errors.New(
	`invalid usage. Expected tomap(iter, func)`)

func tomap(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errToMapUsage
	}

	iter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errToMapUsage
	}

	fn, ok := args[1].(nitro.Callable)
	if !ok {
		return nil, errToMapUsage
	}

	mapv := nitro.NewObject()
	for {
		v, ok, err := nitro.Next(m, iter, 1)
		if err != nil {
			return nil, errToMapUsage
		}
		if !ok {
			break
		}
		res, err := m.Call(fn, []nitro.Value{v[0]}, 1)
		if err != nil {
			return nil, err
		}
		larg, ok := res[0].(*nitro.Array)
		if !ok {
			return nil, fmt.Errorf(
				"conversion func must return a list; instead it returned %v",
				nitro.TypeName(res[0]))
		}
		if larg.Len() != 2 {
			return nil, fmt.Errorf(
				"conversion func must return a list with 2 elements (key and value); "+
					"instead it returned a list with %v elements",
				larg.Len())
		}

		mapv.Put(larg.Get(0), larg.Get(1))
	}

	return []nitro.Value{mapv}, nil
}
