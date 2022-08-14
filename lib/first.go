package lib

import "github.com/dcaiafa/nitro"

func first(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}

	iter, err := getIterArg(vm, args, 0)
	if err != nil {
		return nil, err
	}

	defer vm.IterClose(iter)

	v, err := vm.IterNext(iter, iter.IterNRet())
	if err != nil {
		return nil, err
	}

	if v == nil {
		return make([]nitro.Value, nRet), nil
	}

	return v, nil
}
