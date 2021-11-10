package lib

import "github.com/dcaiafa/nitro"

var errBatchUsage = nitro.NewInvalidUsageError("batch(iter, int)")

func batch(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errBatchUsage
	}

	inIter, err := nitro.MakeIterator(vm, args[0])
	if err != nil {
		return nil, errBatchUsage
	}

	n, ok := args[1].(nitro.Int)
	if !ok {
		return nil, errBatchUsage
	}

	batchIter := &batchIter{
		inIter: inIter,
		n:      int(n.Int64()),
	}

	outIter := nitro.NewIterator(
		batchIter.Next, batchIter.Close, 1)

	return []nitro.Value{outIter}, nil
}

type batchIter struct {
	inIter nitro.Iterator
	n      int
}

func (i *batchIter) Next(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var b = make([]nitro.Value, 0, i.n)
	for len(b) < i.n {
		v, err := vm.IterNext(i.inIter, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		b = append(b, v[0])
	}

	if len(b) == 0 {
		return nil, nil
	}

	return []nitro.Value{nitro.NewArrayFromSlice(b)}, nil
}

func (i *batchIter) Close(vm *nitro.VM) error {
	vm.IterClose(i.inIter)
	return nil
}
