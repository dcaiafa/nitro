package lib

import (
	"github.com/dcaiafa/nitro"
)

func fnFilter(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	e, err := getEnumeratorArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	test, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	iter := &filterIter{
		iter: e,
		test: test,
	}

	closure := nitro.NewClosure(iter.Next, nil)

	return []nitro.Value{closure}, nil
}

type filterIter struct {
	iter nitro.Value
	test nitro.Value
}

func (i *filterIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.iter, 1)
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
