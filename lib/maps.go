package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

func mapsClone(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	m, err := getObjectArg(args, 0)
	if err != nil {
		return nil, err
	}
	r := m.Clone()
	return []nitro.Value{r}, nil
}

func mapsUpdate(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	m, err := getObjectArg(args, 0)
	if err != nil {
		return nil, err
	}

	var u *nitro.Object

	switch arg := args[1].(type) {
	case *nitro.Object:
		u = arg

	case nitro.Callable:
		res, err := vm.Call(arg, []nitro.Value{m}, 1)
		if err != nil {
			return nil, err
		}
		var ok bool
		if u, ok = res[0].(*nitro.Object); !ok {
			return nil, fmt.Errorf(
				"function in argument 2 was expected to return map, but it returned %v",
				nitro.TypeName(res[0]))
		}
	}

	u.ForEach(func(k, v nitro.Value) bool {
		m.Put(k, v)
		return true
	})

	return []nitro.Value{m}, nil
}

func mapsDelete(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	m, err := getObjectArg(args, 0)
	if err != nil {
		return nil, err
	}
	for _, k := range args[1:] {
		m.Delete(k)
	}
	return []nitro.Value{m}, nil
}
