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

	return []nitro.Value{nitro.NewIterator(skipIter.Next, inIter.IterNRet())}, nil
}

type skipIter struct {
	inIter *nitro.Iterator
	skip   int
}

func (i *skipIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.inIter, i.inIter.IterNRet())
		if err != nil {
			return nil, err
		}
		if !ok {
			return []nitro.Value{nitro.NewBool(false), nil}, nil
		}
		if i.skip > 0 {
			i.skip--
			continue
		}
		return append([]nitro.Value{nitro.NewBool(true)}, v...), nil
	}
}
