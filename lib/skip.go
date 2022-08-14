package lib

import (
	"github.com/dcaiafa/nitro"
)

func skip(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
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

	return []nitro.Value{nitro.NewIterator(skipIter.Next, skipIter.Close, inIter.IterNRet())}, nil
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

func (i *skipIter) Close(vm *nitro.VM) error {
	vm.IterClose(i.inIter)
	return nil
}
