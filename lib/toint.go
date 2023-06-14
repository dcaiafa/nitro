package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

func toInt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	if intArg, ok := args[0].(vm.Int); ok {
		return []nitro.Value{intArg}, nil
	}

	inter, ok := args[0].(interface {
		ToInt() (vm.Int, error)
	})
	if !ok {
		return nil, vm.MakeNonRecoverableError(
			fmt.Errorf("%v is not convertible to Int",
				vm.TypeName(args[0])))
	}

	res, err := inter.ToInt()
	if err != nil {
		return nil, err
	}

	return []nitro.Value{res}, nil
}
