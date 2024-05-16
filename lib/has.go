package lib

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

func has(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return []vm.Value{vm.NewBool(false)}, nil
	}

	v := args[1]

	obj, ok := args[0].(interface {
		// TODO: Has should also return an error.
		Has(k vm.Value) bool
	})
	if ok {
		has := obj.Has(v)
		return []vm.Value{vm.NewBool(has)}, nil
	}

	iter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	defer m.IterClose(iter)

	for {
		res, err := m.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return []vm.Value{vm.NewBool(false)}, nil
		}
		if res[0] == v {
			return []vm.Value{vm.NewBool(true)}, nil
		}
	}
}
