package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func has(m *runtime.Machine, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return []runtime.Value{runtime.NewBool(false)}, nil
	}

	obj, ok := args[0].(interface {
		Has(k runtime.Value) bool
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support 'has'",
			runtime.TypeName(args[0]))
	}

	has := obj.Has(args[1])

	return []runtime.Value{runtime.NewBool(has)}, nil
}
