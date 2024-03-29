package lib

import "github.com/dcaiafa/nitro/internal/vm"

func args(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	frameArgs := m.GetCallerArgs()
	argsCopy := make([]vm.Value, len(frameArgs))
	copy(argsCopy, frameArgs)
	res := vm.NewArrayWithSlice(argsCopy)
	return []vm.Value{res}, nil
}
