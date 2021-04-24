package std

import "github.com/dcaiafa/nitro/internal/vm"

func args(m *vm.VM, caps []vm.ValueRef, args []vm.Value, nRet int) ([]vm.Value, error) {
	frameArgs := m.GetArgs()
	argsCopy := make([]vm.Value, len(frameArgs))
	copy(argsCopy, frameArgs)
	res := vm.NewArrayWithSlice(argsCopy)
	return []vm.Value{res}, nil
}
