package lib

import (
	"fmt"

	"github.com/dcaiafa/nitro"
)

type histogramAccum struct {
	hist map[nitro.Value]int
}

func newHistAccum() *histogramAccum {
	return &histogramAccum{
		hist: make(map[nitro.Value]int),
	}
}

func (a *histogramAccum) String() string { return "<histogram>" }
func (a *histogramAccum) Type() string   { return "histogram" }

func (a *histogramAccum) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("histogram does not support this operation")
}

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

var errHistogramUsage = nitro.NewInvalidUsageError(
	"histogram(iter) or histogram(accum, any)")

func hist(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errHistogramUsage
	}

	if len(args) == 2 {
		var accum *histogramAccum
		if args[0] == nil {
			accum = newHistAccum()
		} else {
			var ok bool
			accum, ok = args[0].(*histogramAccum)
			if !ok {
				return nil, errHistogramUsage
			}
		}
		if args[1] == nil {
			return []nitro.Value{accum.ToResult()}, nil
		}

		accum.Process(args[1])
		return []nitro.Value{accum}, nil
	}

	iter, err := nitro.MakeIterator(vm, args[0])
	if err != nil {
		return nil, errHistogramUsage
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
