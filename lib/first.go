package lib

import "github.com/dcaiafa/nitro"

var errFirstUsage = nitro.NewInvalidUsageError("first(iter)")

func first(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errFirstUsage
	}
	iter, err := nitro.MakeIterator(vm, args[0])
	if err != nil {
		return nil, errFirstUsage
	}
	v, err := vm.IterNext(iter, iter.IterNRet())
	if err != nil {
		return nil, err
	}
	if v == nil {
		return make([]nitro.Value, nRet), nil
	}
	return v, nil
}
