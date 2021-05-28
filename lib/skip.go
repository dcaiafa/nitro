package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

func skip(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("not enough arguments")
	}
	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	skip, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	skipIter := &skipIter{inIter: inIter, skip: int(skip)}

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
