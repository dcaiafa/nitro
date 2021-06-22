package lib

import (
	"github.com/dcaiafa/nitro"
)

var errFilterUsage = nitro.NewInvalidUsageError("filter(iter, callable)")

func filter(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errFilterUsage
	}

	inIter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errFilterUsage
	}

	test, ok := args[1].(nitro.Callable)
	if !ok {
		return nil, errFilterUsage
	}

	filterIter := &filterIter{
		inIter: inIter,
		test:   test,
	}

	outIter := nitro.NewIterator(filterIter.Next, nil, inIter.IterNRet())

	return []nitro.Value{outIter}, nil
}

type filterIter struct {
	inIter nitro.Iterator
	test   nitro.Value
}

func (i *filterIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	for {
		v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		res, err := m.Call(i.test, v, 1)
		if err != nil {
			m.IterClose(i.inIter)
			return nil, err
		}
		if nitro.CoerceToBool(res[0]) {
			return v, nil
		}
	}
}
