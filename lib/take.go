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

	return []nitro.Value{nitro.NewIterator(takeIter.Next, nil, inIter.IterNRet())}, nil
}

type takeIter struct {
	inIter nitro.Iterator
	count  int
}

func (i *takeIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.count == 0 {
		return nil, nil
	}

	v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	i.count--
	return v, nil
}
