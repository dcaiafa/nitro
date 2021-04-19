package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func push(m *runtime.Machine, caps []runtime.ValueRef, args []runtime.Value, nRet int) ([]runtime.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	} else if len(args) > 2 {
		return nil, errTooManyArgs
	}

	pushable, ok := args[0].(interface {
		Push(v runtime.Value)
	})
	if !ok {
		return nil, fmt.Errorf(
			"argument %v does not support push",
			runtime.TypeName(args[0]))
	}

	pushable.Push(args[1])
	return nil, nil
}
