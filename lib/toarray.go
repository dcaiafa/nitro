package lib

import (
	"github.com/dcaiafa/nitro"
)

func toarray(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	arr, err := ToArray(m, args[0])
	if err != nil {
		return nil, err
	}

	return []nitro.Value{arr}, nil
}

func ToArray(m *nitro.Machine, v nitro.Value) (*nitro.Array, error) {
	if arr, ok := v.(*nitro.Array); ok {
		return arr, nil
	}

	e, err := nitro.MakeIterator(m, v)
	if err != nil {
		return nil, err
	}

	arr := nitro.NewArray()
	for {
		v, ok, err := nitro.Next(m, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		arr.Push(v[0])
	}

	return arr, nil
}
