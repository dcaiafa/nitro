package lib

import (
	"github.com/dcaiafa/nitro"
)

func fnFilter(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	inEnum, err := getEnumeratorArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	test, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	filterEnum := &filterEnum{
		inEnum: inEnum,
		test:   test,
	}

	outEnum := nitro.NewEnumerator(filterEnum.Next, nil)

	return []nitro.Value{outEnum}, nil
}

type filterEnum struct {
	inEnum nitro.Value
	test   nitro.Value
}

func (i *filterEnum) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.inEnum, 1)
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
