package lib

import "github.com/dcaiafa/nitro"

func mapClone(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
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

func mapUpdate(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
  if len(args) > 2 {
    return nil, errTooManyArgs
  }
  m, err := getObjectArg(args, 0)
  if err != nil {
    return nil, err
  }
  u, err := getObjectArg(args, 1)
  if err != nil {
    return nil, err
  }
  u.ForEach(func(k, v nitro.Value) bool {
    m.Put(k, v)
    return true
  })
  return []nitro.Value{m}, nil
}
  
func mapDelete(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
  m, err := getObjectArg(args, 0)
  if err != nil {
    return nil, err
  }
  for _, k := range args[1:] {
    m.Delete(k)
  }
  return []nitro.Value{m}, nil
}
