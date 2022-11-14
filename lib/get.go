package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

func get(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	} else if len(args) > 2 {
		return nil, errTooManyArgs
	}

	if args[0] == nil {
		return []vm.Value{nil}, nil
	}

	switch t := args[0].(type) {
	case *nitro.Object:
		v, _ := t.Get(args[1])
		return []nitro.Value{v}, nil

	case *nitro.Array:
		i, err := getIntArg(args, 1)
		if err != nil {
			return nil, err
		}
		v := t.Get(int(i))
		return []nitro.Value{v}, nil

	default:
		return nil, fmt.Errorf(
			"cannot get from argument #1 %q", nitro.TypeName(args[0]))
	}
}
