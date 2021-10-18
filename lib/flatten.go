package lib

import "github.com/dcaiafa/nitro"

var errFlattenUsage = nitro.NewInvalidUsageError("flatten(iter)")

func flatten(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	inIter, err := nitro.MakeIterator(vm, args[0])
	if err != nil {
		return nil, errFlattenUsage
	}

	flattenIter := &flattenIter{
		first: inIter,
	}

	outIter := nitro.NewIterator(flattenIter.Next, flattenIter.Close, 1)

	return []nitro.Value{outIter}, nil
}

type flattenIter struct {
	first  nitro.Iterator
	second nitro.Iterator
}

func (i *flattenIter) Next(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.second != nil {
		v, err := vm.IterNext(i.second, 1)
		if err != nil {
			i.Close(vm)
			return nil, err
		}
		if v == nil {
			i.second = nil
			return i.Next(vm, args, nRet)
		}
		return v, nil
	}

	v, err := vm.IterNext(i.first, 1)
	if err != nil {
		i.Close(vm)
		return nil, err
	}

	if v == nil {
		i.Close(vm)
		return nil, nil
	}

	i.second, err = nitro.MakeIterator(vm, v[0])
	if err != nil {
		return v, nil
	}

	return i.Next(vm, args, nRet)
}

func (i *flattenIter) Close(vm *nitro.VM) error {
	vm.IterClose(i.first)
	vm.IterClose(i.second)
	return nil
}
