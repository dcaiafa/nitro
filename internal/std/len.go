package std

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
)

func Len(m *runtime.Machine, caps []runtime.ValueRef, args []runtime.Value, nRet int) ([]runtime.Value, error) {
	if len(args) == 0 {
		return nil, errNotEnoughArgs
	}

	if args[0] == nil {
		return []runtime.Value{runtime.NewInt(0)}, nil
	}

	v, ok := args[0].(interface {
		Len() int
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not have length",
			runtime.TypeName(args[0]))
	}

	return []runtime.Value{runtime.NewInt(int64(v.Len()))}, nil
}
