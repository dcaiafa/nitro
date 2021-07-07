package lib

import (
	"github.com/dcaiafa/nitro"
)

func toarray(m *nitro.VM, args []nitro.Value, _ int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	arr, err := ToArray(m, args[0])
	if err != nil {
		return nil, err
	}

	return []nitro.Value{arr}, nil
}

func ToArray(m *nitro.VM, v nitro.Value) (*nitro.Array, error) {
	if arr, ok := v.(*nitro.Array); ok {
		return arr, nil
	}

	e, err := nitro.MakeIterator(m, v)
	if err != nil {
		return nil, err
	}
	defer m.IterClose(e)

	arr := nitro.NewArray()
	for {
		v, err := m.IterNext(e, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		arr.Push(v[0])
	}

	return arr, nil
}
