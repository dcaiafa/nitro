package lib

import (
	"github.com/dcaiafa/nitro"
)

func filter(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	test, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	filterIter := &filterIter{
		inIter: inIter,
		test:   test,
	}

	outIter := nitro.NewIterator(filterIter.Next, nil, inIter.IterNRet())

	return []nitro.Value{outIter}, nil
}

type filterIter struct {
	inIter *nitro.Iterator
	test   nitro.Value
}

func (i *filterIter) Next(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.inIter, i.inIter.IterNRet())
		if err != nil {
			return nil, err
		}
		if !ok {
			return []nitro.Value{nitro.NewBool(false), nil}, nil
		}
		res, err := m.Call(i.test, v, 1)
		if err != nil {
			return nil, err
		}
		if nitro.CoerceToBool(res[0]) {
			return append([]nitro.Value{nitro.NewBool(true)}, v...), nil
		}
	}
}
