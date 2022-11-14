package lib

import "github.com/dcaiafa/nitro"

func listAppend(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	l, err := getListArg(args, 0)
	if err != nil {
		return nil, err
	}
	for _, arg := range args[1:] {
		l.Add(arg)
	}
	return []nitro.Value{l}, nil
}

func listAppendIter(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	l, err := getListArg(args, 0)
	if err != nil {
		return nil, err
	}
	it, err := getIterArg(m, args, 1)
	if err != nil {
		return nil, err
	}
	for {
		v, err := m.IterNext(it, 1)
		if err != nil {
			return nil, err
		} else if v == nil {
			break
		}
		l.Add(v[0])
	}
	return []nitro.Value{l}, nil
}

func listFind(m *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	l, err := getListArg(args, 0)
	if err != nil {
		return nil, err
	}
	v, err := getValueArg(args, 1)
	if err != nil {
		return nil, err
	}
	ndx := l.Find(v)
	if ndx == -1 {
		return []nitro.Value{nil}, nil
	}
	return []nitro.Value{nitro.NewInt(int64(ndx))}, nil
}
