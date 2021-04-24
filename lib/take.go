package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

func take(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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
	inIter *nitro.Iterator
	count  int
}

func (i *takeIter) Next(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.count == 0 {
		return append(
			[]nitro.Value{nitro.NewBool(false)},
			make([]nitro.Value, nRet-1)...), nil
	}

	v, ok, err := nitro.Next(m, i.inIter, i.inIter.IterNRet())
	if err != nil {
		return nil, err
	}
	if !ok {
		return append(
			[]nitro.Value{nitro.NewBool(false)},
			make([]nitro.Value, nRet-1)...), nil
	}
	i.count--
	return append([]nitro.Value{nitro.NewBool(true)}, v...), nil
}
