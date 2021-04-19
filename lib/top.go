package lib

import "github.com/dcaiafa/nitro"

func top(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	count, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	iter := &topIter{
		inIter: inIter,
		nLeft:  int(count),
	}

	return []nitro.Value{nitro.NewIterator(iter.Next, nil, inIter.IterNRet())}, nil
}

type topIter struct {
	inIter *nitro.Iterator
	nLeft  int
}

func (i *topIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.nLeft <= 0 {
		res := make([]nitro.Value, nRet)
		res[0] = nitro.NewBool(false)
		return res, nil
	}

	v, ok, err := nitro.Next(m, i.inIter, nRet-1)
	if err != nil {
		return nil, err
	}

	if !ok {
		res := make([]nitro.Value, nRet)
		res[0] = nitro.NewBool(false)
		return res, nil
	}

	i.nLeft--

	return append([]nitro.Value{nitro.NewBool(true)}, v...), nil
}
