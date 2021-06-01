package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var errMapReduceUsage = errors.New(
	`invalid usage. Expected mapreduce(iter, string|func, string|func, func+)`)

func mapreduce(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 4 {
		return nil, errMapReduceUsage
	}

	iter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 1: %w", err)
	}
	defer m.IterClose(iter)

	mapExpr, _, err := ParsePathExpr(args[1])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 2: %w", err)
	}

	pickExpr, _, err := ParsePathExpr(args[2])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 3: %w", err)
	}

	reduceFns := make([]nitro.Callable, 0, len(args)-3)
	for i := 3; i < len(args); i++ {
		reduceFn, ok := args[i].(nitro.Callable)
		if !ok {
			return nil, errMapReduceUsage
		}
		reduceFns = append(reduceFns, reduceFn)
	}

	res := nitro.NewObject()
	for {
		cur, err := m.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if cur == nil {
			break
		}

		mapKey, err := mapExpr.Eval(m, cur[0])
		if err != nil {
			return nil, err
		}
		if mapKey == nil {
			continue
		}

		val, err := pickExpr.Eval(m, cur[0])
		if err != nil {
			return nil, err
		}
		if val == nil {
			continue
		}

		var accumList *nitro.Array
		accumRef, _ := res.IndexRef(mapKey)
		if *accumRef.Refo() == nil {
			accumList = nitro.NewArrayFromSlice(make([]nitro.Value, len(reduceFns)))
			*accumRef.Refo() = accumList
		} else {
			accumList = (*accumRef.Refo()).(*nitro.Array)
		}

		for i, reduceFn := range reduceFns {
			partialRes, err := m.Call(
				reduceFn, []nitro.Value{accumList.Get(i), val}, 1)
			if err != nil {
				return nil, err
			}
			accumList.Put(i, partialRes[0])
		}
	}

	err = nil
	res.ForEach(func(k, accum nitro.Value) bool {
		accumList := accum.(*nitro.Array)
		for i, reduceFn := range reduceFns {
			finalRes, callErr := m.Call(reduceFn, []nitro.Value{accumList.Get(i)}, 1)
			if callErr != nil {
				err = callErr
				return false
			}
			accumList.Put(i, finalRes[0])
		}
		if len(reduceFns) == 1 {
			res.Put(k, accumList.Get(0))
		}
		return true
	})

	if err != nil {
		return nil, fmt.Errorf("final reduce failed: %w", err)
	}

	return []nitro.Value{res}, nil
}
