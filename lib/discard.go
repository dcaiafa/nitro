package lib

import (
	"io"
	"io/ioutil"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

func discard(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	} else if len(args) > 1 {
		return nil, errTooManyArgs
	}

	if reader, ok := args[0].(io.Reader); ok {
		_, err := io.Copy(ioutil.Discard, reader)
		if err != nil {
			return nil, err
		}
		core.CloseReader(reader)
	} else if iter, err := nitro.MakeIterator(vm, args[0]); err == nil {
		for {
			v, err := vm.IterNext(iter, 1)
			if err != nil {
				return nil, err
			}
			if v == nil {
				break
			}
		}
	} else {
		return nil, errExpectedArg(0, args[0], "reader", "iter")
	}

	return nil, nil
}
