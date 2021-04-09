package lib

import (
	"github.com/dcaiafa/nitro"
)

func fnFilter(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

	outIter := nitro.NewIterator(filterIter.Next, nil)

	return []nitro.Value{outIter}, nil
}

type filterIter struct {
	inIter nitro.Value
	test   nitro.Value
}

func (i *filterIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.inIter, 1)
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
			return []nitro.Value{nitro.NewBool(true), v[0]}, nil
		}
	}
}
