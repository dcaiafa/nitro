package lib

import (
	"github.com/dcaiafa/nitro"
)

var errMapUsage = nitro.NewInvalidUsageError("map(iter, callable)")

func mapp(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	inIter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errMapUsage
	}
	fn, ok := args[1].(nitro.Callable)
	if !ok {
		return nil, errMapUsage
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
	v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
	if err != nil {
		i.Close(m)
		return nil, err
	}
	if v == nil {
		i.Close(m)
		return nil, nil
	}
	res, err := m.Call(i.fn, v, 1)
	if err != nil {
		i.Close(m)
		return nil, err
	}

	return []nitro.Value{res[0]}, nil
}

func (i *mapIter) Close(m *nitro.VM) error {
	m.IterClose(i.inIter)
	return nil
}
