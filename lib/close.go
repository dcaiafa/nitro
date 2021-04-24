package lib

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

func closep(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	f, ok := args[0].(io.Closer)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument 1 to be File, but it is %v",
			nitro.TypeName(args[0]))
	}

	err := f.Close()
	if err != nil && !errors.Is(err, os.ErrClosed) {
		return nil, err
	}

	return nil, nil
}
