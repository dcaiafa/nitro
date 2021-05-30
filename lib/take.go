package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

func take(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("not enough arguments")
	}
	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	count, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	takeIter := &takeIter{inIter: inIter, count: int(count)}

	return []nitro.Value{nitro.NewIterator(takeIter.Next, takeIter.Close, inIter.IterNRet())}, nil
}

type takeIter struct {
	inIter nitro.Iterator
	count  int
}

func (i *takeIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.count == 0 {
		i.Close(m)
		return nil, nil
	}

	v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
	if err != nil {
		i.Close(m)
		return nil, err
	}
	if v == nil {
		i.Close(m)
		return nil, nil
	}
	i.count--
	return v, nil
}

func (i *takeIter) Close(m *nitro.VM) error {
	m.IterClose(i.inIter)
	return nil
}
