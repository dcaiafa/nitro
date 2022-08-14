package lib

import "github.com/dcaiafa/nitro"

func dop(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	iter, err := getIterArg(vm, args, 0)
	if err != nil {
		return nil, err
	}

	fn, err := getCallableArg(args, 1)
	if err != nil {
		return nil, err
	}

	for {
		v, err := vm.IterNext(iter, iter.IterNRet())
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		_, err = vm.Call(fn, v, 0)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
