package lib

import (
	"github.com/dcaiafa/nitro"
)

func mapp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	fn, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	mapIter := &mapIter{
		inIter: inIter,
		fn:     fn,
	}

	outIter := nitro.NewIterator(mapIter.Next, nil, 1)
	return []nitro.Value{outIter}, nil
}

type mapIter struct {
	inIter nitro.Iterator
	fn     nitro.Value
}

func (i *mapIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	v, ok, err := nitro.Next(m, i.inIter, i.inIter.IterNRet())
	if err != nil {
		return nil, err
	}
	if !ok {
		return iterDone(nRet)
	}
	res, err := m.Call(i.fn, v, 1)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewBool(true), res[0]}, nil
}
