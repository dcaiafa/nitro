package vm

import (
	"fmt"
)

type Iterable interface {
	Value
	MakeIterator() Iterator
}

func MakeIterator(m *VM, v Value) (Iterator, error) {
	if v == nil {
		return NewIterator(emptyIter, nil, 1), nil
	}
	switch v := v.(type) {
	case Iterator:
		return v, nil
	case Iterable:
		return v.MakeIterator(), nil
	default:
		return nil, fmt.Errorf("Value of type %q %w", TypeName(v), ErrIsNotIterable)
	}
}

func emptyIter(m *VM, args []Value, nret int) ([]Value, error) {
	if nret < 2 {
		nret = 2
	}
	res := make([]Value, nret)
	res[0] = False
	return res, nil
}
