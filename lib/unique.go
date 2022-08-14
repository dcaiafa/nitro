package lib

import (
	"github.com/dcaiafa/nitro"
)

func unique(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 2); err != nil {
		return nil, err
	}

	e, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	var keyFn *PathExpr
	if len(args) == 2 {
		var err error
		keyFn, _, err = ParsePathExpr(args[1])
		if err != nil {
			return nil, err
		}
	}

	set := make(map[nitro.Value]nitro.Value)

	for {
		v, err := m.IterNext(e, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		key := v[0]
		if keyFn != nil {
			key, err = keyFn.Eval(m, v[0])
			if err != nil {
				return nil, err
			}
		}

		set[key] = v[0]
	}

	arr := nitro.NewArrayFromSlice(make([]nitro.Value, 0, len(set)))
	for _, v := range set {
		arr.Add(v)
	}

	return []nitro.Value{arr}, nil
}
