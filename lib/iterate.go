package lib

import "github.com/dcaiafa/nitro"

func iterate(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
  if err := expectArgCount(args, 1, 1); err != nil {
    return nil, err
  }

  iter, err := getIterArg(vm, args, 0)
  if err != nil {
    return nil, err
  }

  return []nitro.Value{iter}, nil
}
