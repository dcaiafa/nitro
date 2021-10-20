package lib

import (
	"io"
	"io/ioutil"

	"github.com/dcaiafa/nitro"
)

var errDiscardUsage = nitro.NewInvalidUsageError("discard(iter|reader)")

func discard(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errDiscardUsage
	}

	if reader, ok := args[0].(io.Reader); ok {
		_, err := io.Copy(ioutil.Discard, reader)
		if err != nil {
			return nil, err
		}
		CloseReader(reader)
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
		return nil, errDiscardUsage
	}

	return nil, nil
}
