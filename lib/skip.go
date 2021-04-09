package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

func fnSkip(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

	return []nitro.Value{nitro.NewIterator(skipIter.Next, nil)}, nil
}

type skipIter struct {
	inIter nitro.Value
	skip   int
}

func (i *skipIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	for {
		v, ok, err := nitro.Next(m, i.inIter, 1)
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
		return []nitro.Value{nitro.NewBool(true), v[0]}, nil
	}
}
