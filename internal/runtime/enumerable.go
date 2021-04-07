package runtime

import (
	"fmt"
	"log"
)

type Enumerable interface {
	Value
	Enumerate() *Closure
}

func MakeEnumerator(m *Machine, v Value) (Value, error) {
	switch v := v.(type) {
	case *Closure:
		return v, nil
	case Enumerable:
		return v.Enumerate(), nil
	default:
		return nil, fmt.Errorf("Value of type %q %w", v.Type(), ErrIsNotEnumerable)
	}
}

func Next(m *Machine, e Value, n int) ([]Value, bool, error) {
	c, ok := e.(*Closure)
	if !ok {
		return nil, false, fmt.Errorf("not an enumerator")
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
			"enumerator's first return value must be a Bool, but instead it was %q",
			ret[0].Type())
	}

	if !hasValue.Bool() {
		return nil, false, nil
	}

	return ret[1:], true, nil
}
