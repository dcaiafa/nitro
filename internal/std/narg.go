package std

import "github.com/dcaiafa/nitro/internal/runtime"

func narg(m *runtime.Machine, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	return []runtime.Value{runtime.NewInt(int64(m.GetNArg()))}, nil
}
