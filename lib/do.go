package lib

import "github.com/dcaiafa/nitro"

var errDoUsage = nitro.NewInvalidUsageError("do(iter, fn)")

func dop(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errDoUsage
	}

	iter, err := nitro.MakeIterator(vm, args[0])
	if err != nil {
		return nil, errDoUsage
	}

	fn, ok := args[1].(nitro.Callable)
	if !ok {
		return nil, errDoUsage
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
