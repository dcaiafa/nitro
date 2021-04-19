package lib

import "github.com/dcaiafa/nitro"

func mapreduce(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	iter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	mapFn, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	reducePickFn, err := getCallableArg(args, 2)
	if err != nil {
		return nil, err
	}

	reduceFn, err := getCallableArg(args, 3)
	if err != nil {
		return nil, err
	}

	res := nitro.NewObject()
	for {
		cur, has, err := nitro.Next(m, iter, 1)
		if err != nil {
			return nil, err
		}
		if !has {
			break
		}
		mapKey, err := m.Call(mapFn, []nitro.Value{cur[0]}, 1)
		if err != nil {
			return nil, err
		}

		val, err := m.Call(reducePickFn, []nitro.Value{cur[0]}, 1)
		if err != nil {
			return nil, err
		}

		if val[0] == nil {
			continue
		}

		accumRef, _ := res.IndexRef(mapKey[0])
		accumRes, err := m.Call(
			reduceFn, []nitro.Value{*accumRef.Refo(), val[0]}, 1)
		if err != nil {
			return nil, err
		}

		*accumRef.Refo() = accumRes[0]
	}

	err = nil
	res.ForEach(func(k, accum nitro.Value) bool {
		accumRes, callErr := m.Call(reduceFn, []nitro.Value{accum}, 1)
		if callErr != nil {
			err = callErr
			return false
		}
		res.Put(k, accumRes[0])
		return true
	})

	return []nitro.Value{res}, nil
}
