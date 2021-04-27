package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

func push(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	} else if len(args) > 2 {
		return nil, errTooManyArgs
	}

	pushable, ok := args[0].(interface {
		Push(v vm.Value)
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support push",
			vm.TypeName(args[0]))
	}

	pushable.Push(args[1])
	return nil, nil
}
