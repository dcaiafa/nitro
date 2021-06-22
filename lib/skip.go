package lib

import (
	"github.com/dcaiafa/nitro"
)

var errSkipUsage = nitro.NewInvalidUsageError("skip(iter, int)")

func skip(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errSkipUsage
	}
	inIter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errSkipUsage
	}
	skip, ok := args[1].(nitro.Int)
	if !ok {
		return nil, errSkipUsage
	}

	skipIter := &skipIter{inIter: inIter, skip: int(skip.Int64())}

	return []nitro.Value{nitro.NewIterator(skipIter.Next, nil, inIter.IterNRet())}, nil
}

type skipIter struct {
	inIter nitro.Iterator
	skip   int
}

func (i *skipIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	for {
		v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		if i.skip > 0 {
			i.skip--
			continue
		}
		return v, nil
	}
}
