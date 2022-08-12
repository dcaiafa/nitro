package lib

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

func closep(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	} else if len(args) > 1 {
		return nil, errTooManyArgs
	}

	switch arg := args[0].(type) {
	case io.Closer:
		err := arg.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			return nil, err
		}

	case nitro.Iterator:
		err := m.IterClose(arg)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf(
			"arg #1 %v is not closeable",
			nitro.TypeName(args[0]))
	}

	return nil, nil
}
