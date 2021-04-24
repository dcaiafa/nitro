package std

import "github.com/dcaiafa/nitro/internal/vm"

func narg(m *vm.Machine, caps []vm.ValueRef, args []vm.Value, nRet int) ([]vm.Value, error) {
	return []vm.Value{vm.NewInt(int64(m.GetNArg()))}, nil
}
