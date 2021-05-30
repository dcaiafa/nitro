package lib

import (
	"errors"

	"github.com/dcaiafa/nitro"
)

var errReduceUsage = errors.New(
	`invalid usage. Expected reduce(iter, func)`)

func reduce(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errReduceUsage
	}

	iter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errReduceUsage
	}
	defer m.IterClose(iter)

	reducer, ok := args[1].(nitro.Callable)
	if !ok {
		return nil, errReduceUsage
	}

	var accum nitro.Value
	for {
		val, err := m.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if val == nil {
			break
		}
		res, err := m.Call(reducer, []nitro.Value{accum, val[0]}, 1)
		if err != nil {
			return nil, err
		}
		accum = res[0]
	}

	res, err := m.Call(reducer, []nitro.Value{accum}, 1)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{res[0]}, nil
}
