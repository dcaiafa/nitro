package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

func delete_(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return nil, nil
	}

	obj, ok := args[0].(interface {
		Delete(key vm.Value)
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support 'delete'",
			vm.TypeName(args[0]))
	}

	obj.Delete(args[1])
	return nil, nil
}
