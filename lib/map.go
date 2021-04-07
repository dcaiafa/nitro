package lib

import (
	"github.com/dcaiafa/nitro"
)

func fnMap(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	inEnum, err := getEnumeratorArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	fn, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	mapEnum := &mapEnum{
		inEnum: inEnum,
		fn:     fn,
	}

	outEnum := nitro.NewEnumerator(mapEnum.Next, nil)
	return []nitro.Value{outEnum}, nil
}

type mapEnum struct {
	inEnum nitro.Value
	fn     nitro.Value
}

func (i *mapEnum) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	v, ok, err := nitro.Next(m, i.inEnum, 1)
	if err != nil {
		return nil, err
	}
	if !ok {
		return []nitro.Value{nitro.NewBool(false), nil}, nil
	}
	res, err := m.Call(i.fn, v, 1)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewBool(true), res[0]}, nil
}
