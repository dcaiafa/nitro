package std

import "github.com/dcaiafa/nitro/internal/runtime"

func args(m *runtime.Machine, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	frameArgs := m.GetArgs()
	argsCopy := make([]runtime.Value, len(frameArgs))
	copy(argsCopy, frameArgs)
	res := runtime.NewArrayWithSlice(argsCopy)
	return []runtime.Value{res}, nil
}
