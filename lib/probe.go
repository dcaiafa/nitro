package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
	libio "github.com/dcaiafa/nitro/lib/io"
)

func probe(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if nret != 1 {
		return nil, fmt.Errorf(
			"probe must be used as an intermediate between two functions")
	}

	probeValue := args[0]
	otherArgs := make([]nitro.Value, len(args)-1)
	copy(otherArgs, args[1:])

	if inIter, ok := probeValue.(nitro.Iterator); ok {
		probeIter := &probeIter{
			inIter:    inIter,
			otherArgs: otherArgs,
		}
		return []nitro.Value{
			nitro.NewIterator(probeIter.Next, probeIter.Close, inIter.IterNRet())}, nil
	}

	otherArgs = append(otherArgs, probeValue)
	_, err := basePrint(libio.Stderr(vm), vm, otherArgs, 0)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{probeValue}, nil
}

type probeIter struct {
	inIter    nitro.Iterator
	otherArgs []nitro.Value
}

func (i *probeIter) Next(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	v, err := vm.IterNext(i.inIter, i.inIter.IterNRet())
	if err != nil {
		i.Close(vm)
		return nil, err
	}
	if v == nil {
		i.Close(vm)
		return nil, nil
	}
	printArgs := append(i.otherArgs, v...)
	basePrint(libio.Stderr(vm), vm, printArgs, 0)

	return v, nil
}

func (i *probeIter) Close(vm *nitro.VM) error {
	vm.IterClose(i.inIter)
	return nil
}
