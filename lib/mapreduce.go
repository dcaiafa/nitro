package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var errMapReduceUsage = errors.New(
	`invalid usage. Expected mapreduce(iter, string|func, string|func, func)`)

func mapreduce(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 4 {
		return nil, errMapReduceUsage
	}

	iter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 1: %w", err)
	}

	mapExpr, _, err := ParsePathExpr(args[1])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 2: %w", err)
	}

	pickExpr, _, err := ParsePathExpr(args[2])
	if err != nil {
		return nil, fmt.Errorf("invalid argument 3: %w", err)
	}

	reduceFn, ok := args[3].(nitro.Callable)
	if !ok {
		return nil, errMapReduceUsage
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

		accumRef, _ := res.IndexRef(mapKey)
		accumRes, err := m.Call(
			reduceFn, []nitro.Value{*accumRef.Refo(), val}, 1)
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

	if err != nil {
		return nil, fmt.Errorf("final reduce failed: %w", err)
	}

	return []nitro.Value{res}, nil
}
