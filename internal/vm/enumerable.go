package vm

import (
	"fmt"
	"log"
)

type Iterable interface {
	Value
	Iterate() *Iterator
}

func MakeIterator(m *Machine, v Value) (*Iterator, error) {
	switch v := v.(type) {
	case *Iterator:
		return v, nil
	case Iterable:
		return v.Iterate(), nil
	default:
		return nil, fmt.Errorf("Value of type %q %w", TypeName(v), ErrIsNotIterable)
	}
}

func Next(m *Machine, e Value, n int) ([]Value, bool, error) {
	c, ok := e.(*Iterator)
	if !ok {
		return nil, false, fmt.Errorf("not an iterator")
	}

	if n == 0 {
		log.Panic("n cannot be zero")
	}

	ret, err := m.Call(c, nil, n+1)
	if err != nil {
		return nil, false, err
	}

	hasValue, ok := ret[0].(Bool)
	if !ok {
		return nil, false, fmt.Errorf(
			"iterator's first return value must be a Bool, but instead it was %q",
			TypeName(ret[0]))
	}

	if !hasValue.Bool() {
		return nil, false, nil
	}

	return ret[1:], true, nil
}
