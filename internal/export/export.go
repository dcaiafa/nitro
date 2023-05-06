package export

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

type Type int

const (
	Func Type = iota
	Value
)

type Export struct {
	N string
	T Type
	F func(vm *vm.VM, args []vm.Value, nret int) ([]vm.Value, error)
	V vm.Value
}

func (e *Export) Value() vm.Value {
	switch e.T {
	case Func:
		return vm.NewNativeFn(e.F)
	case Value:
		return e.V
	default:
		panic("not reached")
	}
}

type Exports []Export
