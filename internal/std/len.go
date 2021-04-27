package std

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
)

func Len(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if len(args) == 0 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return []vm.Value{vm.NewInt(0)}, nil
	}

	v, ok := args[0].(interface {
		Len() int
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not have length",
			vm.TypeName(args[0]))
	}

	return []vm.Value{vm.NewInt(int64(v.Len()))}, nil
}
