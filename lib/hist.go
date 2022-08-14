package lib

import (
	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

type histogramAccum struct {
	hist map[nitro.Value]int
}

func newHistAccum() *histogramAccum {
	return &histogramAccum{
		hist: make(map[nitro.Value]int),
	}
}

func (a *histogramAccum) String() string    { return "<histogram>" }
func (a *histogramAccum) Type() string      { return "histogram" }
func (a *histogramAccum) Traits() vm.Traits { return vm.TraitNone }

func (a *histogramAccum) Process(v nitro.Value) {
	a.hist[v]++
}

func (a *histogramAccum) ToResult() *nitro.Object {
	r := nitro.NewObject()
	for k, v := range a.hist {
		r.Put(k, nitro.NewInt(int64(v)))
	}
	return r
}

func hist(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	} else if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	if len(args) == 2 {
		var accum *histogramAccum
		if args[0] == nil {
			accum = newHistAccum()
		} else {
			var ok bool
			accum, ok = args[0].(*histogramAccum)
			if !ok {
				return nil, errExpectedArg(0, args[0], "hist")
			}
		}

		if args[1] == nil {
			return []nitro.Value{accum.ToResult()}, nil
		}

		accum.Process(args[1])
		return []nitro.Value{accum}, nil
	}

	iter, err := getIterArg(vm, args, 0)
	if err != nil {
		return nil, err
	}

	accum := newHistAccum()

	for {
		v, err := vm.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		accum.Process(v[0])
	}

	return []nitro.Value{accum.ToResult()}, nil
}
