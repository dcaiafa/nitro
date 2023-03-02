package export

import (
	"github.com/dcaiafa/nitro/internal/vm"
)

type Type int

const (
	Func Type = iota
	String
	Int
	Float
	Custom
)

type Export struct {
	N  string
	T  Type
	F  func(vm *vm.VM, args []vm.Value, nret int) ([]vm.Value, error)
	S  string
	I  int64
	Fl float64
	C  func(e *Export) vm.Value
}

func (e *Export) Value() vm.Value {
	switch e.T {
	case Func:
		return vm.NewNativeFn(e.F)
	case String:
		return vm.NewString(e.S)
	case Int:
		return vm.NewInt(e.I)
	case Float:
		return vm.NewFloat(e.Fl)
	case Custom:
		return e.C(e)
	default:
		panic("not reached")
	}
}

type Exports []Export
