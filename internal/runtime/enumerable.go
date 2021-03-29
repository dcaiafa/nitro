package runtime

import (
	"context"
	"fmt"
	"log"
)

type Enumerable interface {
	Value
	Enumerate() *Closure
}

func MakeEnumerator(ctx context.Context, v Value) (Value, error) {
	switch v := v.(type) {
	case *Closure:
		return v, nil
	case Enumerable:
		return v.Enumerate(), nil
	default:
		return nil, fmt.Errorf("Value of type %q %w", v.Type(), ErrIsNotEnumerable)
	}
}

func Next(ctx context.Context, e Value, n int) ([]Value, bool, error) {
	c, ok := e.(*Closure)
	if !ok {
		return nil, false, fmt.Errorf("not an enumerator")
	}

	if n == 0 {
		log.Panic("n cannot be zero")
	}

	ret, err := Call(ctx, c, nil, n+1)
	if err != nil {
		return nil, false, err
	}

	hasValue, ok := ret[len(ret)-1].(Bool)
	if !ok {
		return nil, false, fmt.Errorf(
			"enumerator's last return value must be a Bool, but instead it was %q",
			ret[len(ret)-1].Type())
	}

	if !hasValue.Bool() {
		return nil, false, nil
	}

	return ret[:len(ret)-1], true, nil
}

func Call(ctx context.Context, callable Value, args []Value, retN int) ([]Value, error) {
	m := MachineFromContext(ctx)
	return m.runCallable(ctx, callable, args, retN)
}
