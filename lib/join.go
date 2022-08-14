package lib

import (
	"strings"

	"github.com/dcaiafa/nitro"
)

func join(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	iter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	list, err := ToArray(m, iter)
	if err != nil {
		return nil, err
	}

	sep, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	elems := make([]string, list.Len())
	for i := 0; i < len(elems); i++ {
		elems[i] = list.Get(i).String()
	}

	res := strings.Join(elems, sep)
	return []nitro.Value{nitro.NewString(res)}, nil
}
