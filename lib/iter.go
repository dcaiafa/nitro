package lib

import "github.com/dcaiafa/nitro"

func iter_into(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
  if err := expectArgCount(args, 1, 1); err != nil {
    return nil, err
  }

  iter, err := getIterArg(vm, args, 0)
  if err != nil {
    return nil, err
  }

  return []nitro.Value{iter}, nil
}

func iter_next(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
  iter, err := getIterArg(vm, args, 0)
  if err != nil {
    return nil, err
  }
  res, err := vm.IterNext(iter, nret)
  if err != nil {
    return nil, err
  }
  if res == nil {
    vm.IterClose(iter)
    return make([]nitro.Value, nret), nil
  }
  return res, nil
}

