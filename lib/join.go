package lib

import (
	"fmt"
	"strings"

	"github.com/dcaiafa/nitro"
)

var errJoinUsage = nitro.NewInvalidUsageError("join(iter, string)")

func join(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errJoinUsage
	}

	list, err := ToArray(m, args[0])
	if err != nil {
		return nil, fmt.Errorf(
			"%w: argument 1 is not an array or iterable", errJoinUsage)
	}

	sep, ok := args[1].(nitro.String)
	if !ok {
		return nil, errJoinUsage
	}

	elems := make([]string, list.Len())
	for i := 0; i < len(elems); i++ {
		elems[i] = list.Get(i).String()
	}

	res := strings.Join(elems, sep.String())

	return []nitro.Value{nitro.NewString(res)}, nil
}
