package lib

import "github.com/dcaiafa/nitro"

func mapreduce(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	iter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	mapFn, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	reduceFn, err := getCallableArg(args, 2)
	if err != nil {
		return nil, err
	}

	res := nitro.NewObject()
	for {
		v, ok, err := nitro.Next(m, iter, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}

		mapKey, err := m.Call(mapFn, []nitro.Value{v[0]}, 1)
		if err != nil {
			return nil, err
		}

		accumRef, _ := res.IndexRef(mapKey[0])
		accumRes, err := m.Call(reduceFn, []nitro.Value{*accumRef.Refo()}, 1)
		if err != nil {
			return nil, err
		}

		*accumRef.Refo() = accumRes[0]
	}

	return []nitro.Value{res}, nil
}
