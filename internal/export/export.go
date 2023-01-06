package export

import (
	gosort "sort"

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
	P  string
	N  string
	T  Type
	F  func(vm *vm.VM, args []vm.Value, nret int) ([]vm.Value, error)
	S  string
	I  int64
	Fl float64
	C  func(e *Export) vm.Value
}

type Exports []Export

func (e Exports) GetPackageExports(pkg string, handleSym func(symName string, value vm.Value)) {
	i := gosort.Search(len(e), func(i int) bool { return e[i].P >= pkg })

	for ; i < len(e) && e[i].P == pkg; i++ {
		handleSym(e[i].N, e.getExportValue(i))
	}
}

func (e Exports) getExportValue(index int) vm.Value {
	switch e[index].T {
	case Func:
		return vm.NewNativeFn(e[index].F)
	case String:
		return vm.NewString(e[index].S)
	case Int:
		return vm.NewInt(e[index].I)
	case Float:
		return vm.NewFloat(e[index].Fl)
	case Custom:
		return e[index].C(&e[index])
	default:
		panic("not reached")
	}
}

func (e Exports) Len() int {
	return len(e)
}

func (e Exports) Less(i, j int) bool {
	if e[i].P < e[j].P {
		return true
	} else if e[i].P > e[j].P {
		return false
	} else {
		return e[i].N < e[j].N
	}
}

func (e *Exports) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}
