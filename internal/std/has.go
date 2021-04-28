package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

func has(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return []vm.Value{vm.NewBool(false)}, nil
	}

	obj, ok := args[0].(interface {
		// TODO: Has should also return an error.
		Has(k vm.Value) bool
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support 'has'",
			vm.TypeName(args[0]))
	}

	has := obj.Has(args[1])

	return []vm.Value{vm.NewBool(has)}, nil
}
