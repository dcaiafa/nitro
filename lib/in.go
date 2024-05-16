package lib

import "github.com/dcaiafa/nitro"

func in(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	v := args[0]

	found := false
	for _, arg := range args[1:] {
		if v == arg {
			found = true
			break
		}
	}

	return []nitro.Value{nitro.NewBool(found)}, nil
}
