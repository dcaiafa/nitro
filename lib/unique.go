package lib

import (
	"github.com/dcaiafa/nitro"
)

var errUniqueUsage = nitro.NewInvalidUsageError("unique(iter)")

func unique(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errUniqueUsage
	}

	e, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errUniqueUsage
	}

	set := make(map[nitro.Value]bool)

	for {
		v, err := m.IterNext(e, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		set[v[0]] = true
	}

	arr := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(set)))
	for v := range set {
		arr.Add(v)
	}

	return []nitro.Value{arr}, nil
}
