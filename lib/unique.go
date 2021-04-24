package lib

import (
	"github.com/dcaiafa/nitro"
)

func unique(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	e, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	set := make(map[nitro.Value]bool)

	for {
		v, ok, err := nitro.Next(m, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		set[v[0]] = true
	}

	arr := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(set)))
	for v := range set {
		arr.Push(v)
	}

	return []nitro.Value{arr}, nil
}
