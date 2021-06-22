package lib

import (
	"github.com/dcaiafa/nitro"
)

var errTakeUsage = nitro.NewInvalidUsageError("take(iter, int)")

func take(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errTakeUsage
	}

	inIter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errTakeUsage
	}

	count, ok := args[1].(nitro.Int)
	if !ok {
		return nil, errTakeUsage
	}

	takeIter := &takeIter{inIter: inIter, count: int(count.Int64())}

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
