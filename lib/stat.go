package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

type avgAccum struct {
	sum   float64
	count int
}

func (a *avgAccum) String() string    { return "<avg>" }
func (a *avgAccum) Type() string      { return "avg" }
func (a *avgAccum) Traits() vm.Traits { return vm.TraitNone }

var errAvgUsage = errors.New(
	"invalid usage. Expected: avg(iter) or avg(accum, int|float?)")

func avg(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errAvgUsage
	}

	if len(args) == 1 && args[0] == nil {
		// Special case: avg(nil) == nil
		// so that reduce([], avg) == nil.
		return []nitro.Value{nil}, nil
	}

	// Form 1: avg(accum|nil, float|int|nil?)
	accum, ok := args[0].(*avgAccum)
	if ok || args[0] == nil {
		if len(args) != 1 && len(args) != 2 {
			return nil, errAvgUsage
		}
		if accum == nil {
			accum = new(avgAccum)
		}
		if len(args) == 1 || args[1] == nil {
			if accum.count <= 0 {
				return nil, errAvgUsage
			}
			return []nitro.Value{nitro.NewFloat(accum.sum / float64(accum.count))}, nil
		}
		accum.count++
		switch v := args[1].(type) {
		case nitro.Int:
			accum.sum += float64(v.Int64())
		case nitro.Float:
			accum.sum += v.Float64()
		default:
			return nil, fmt.Errorf(
				"%w. Argument 2 was %v",
				errAvgUsage, nitro.TypeName(v))
		}
		return []nitro.Value{accum}, nil
	}

	// Form 2: avg(float|int...)
	switch args[0].(type) {
	case nitro.Int, nitro.Float:
		var sum float64
		for i, arg := range args {
			switch v := arg.(type) {
			case nitro.Int:
				sum += float64(v.Int64())
			case nitro.Float:
				sum += v.Float64()
			default:
				return nil, fmt.Errorf(
					"%w. Argument %d was %v",
					errAvgUsage, i, nitro.TypeName(v))
			}
		}
		return []nitro.Value{nitro.NewFloat(sum / float64(len(args)))}, nil
	}

	// Form 3: avg(iter[float|int])
	iter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errAvgUsage
	}

	count := 0
	var sum float64
	for {
		v, err := m.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if v == nil {
			break
		}
		count++
		switch v := v[0].(type) {
		case nitro.Int:
			sum += float64(v.Int64())
		case nitro.Float:
			sum += v.Float64()
		default:
			return nil, fmt.Errorf(
				"%w. iterator returned %v",
				errAvgUsage, nitro.TypeName(v))
		}
	}

	if count == 0 {
		return []nitro.Value{nil}, nil
	}

	return []nitro.Value{nitro.NewFloat(sum / float64(count))}, nil
}

var errMaxUsage = errors.New(
	"invalid usage. Expected: max(int|float...) or max(iter)")

func max(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errMaxUsage
	}

	if args[0] != nil {
		iter, err := nitro.MakeIterator(m, args[0])
		if err == nil {
			var maxV nitro.Value
			for {
				v, err := m.IterNext(iter, 1)
				if err != nil {
					return nil, err
				}
				if v == nil {
					break
				}
				if v[0] == nil {
					continue
				}
				if maxV == nil {
					maxV = v[0]
					continue
				}
				isGT, err := evalCmpOp(nitro.OpGT, v[0], maxV)
				if err != nil {
					return nil, err
				}
				if isGT {
					maxV = v[0]
				}
			}
			return []nitro.Value{maxV}, nil
		}
	}

	var maxV nitro.Value
	for _, arg := range args {
		if arg == nil {
			continue
		}
		if maxV == nil {
			maxV = arg
			continue
		}
		isGT, err := evalCmpOp(nitro.OpGT, arg, maxV)
		if err != nil {
			return nil, err
		}
		if isGT {
			maxV = arg
		}
	}
	return []nitro.Value{maxV}, nil
}

var errMinUsage = errors.New(
	"invalid usage. Expected: min(int|float...) or min(iter)")

func min(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errMinUsage
	}

	if args[0] != nil {
		iter, err := nitro.MakeIterator(m, args[0])
		if err == nil {
			var minV nitro.Value
			for {
				v, err := m.IterNext(iter, 1)
				if err != nil {
					return nil, err
				}
				if v == nil {
					break
				}
				if v[0] == nil {
					continue
				}
				if minV == nil {
					minV = v[0]
					continue
				}
				isGT, err := evalCmpOp(nitro.OpLT, v[0], minV)
				if err != nil {
					return nil, err
				}
				if isGT {
					minV = v[0]
				}
			}
			return []nitro.Value{minV}, nil
		}
	}

	var minV nitro.Value
	for _, arg := range args {
		if arg == nil {
			continue
		}
		if minV == nil {
			minV = arg
			continue
		}
		isGT, err := evalCmpOp(nitro.OpLT, arg, minV)
		if err != nil {
			return nil, err
		}
		if isGT {
			minV = arg
		}
	}
	return []nitro.Value{minV}, nil
}

var errSumUsage = errors.New(
	"invalid usage. Expected: sum(int|float...) or sum(iter)")

func sum(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error

	if len(args) == 0 {
		return nil, errMinUsage
	}

	if args[0] != nil {
		iter, err := nitro.MakeIterator(m, args[0])
		if err == nil {
			var sumV nitro.Value
			for {
				v, err := m.IterNext(iter, 1)
				if err != nil {
					return nil, err
				}
				if v == nil {
					break
				}
				if v[0] == nil {
					continue
				}
				if sumV == nil {
					sumV = v[0]
					continue
				}
				sumV, err = nitro.EvalOp(nitro.OpAdd, sumV, v[0])
				if err != nil {
					return nil, err
				}
			}
			return []nitro.Value{sumV}, nil
		}
	}

	var sumV nitro.Value
	for _, arg := range args {
		if arg == nil {
			continue
		}
		if sumV == nil {
			sumV = arg
			continue
		}
		sumV, err = nitro.EvalOp(nitro.OpAdd, sumV, arg)
		if err != nil {
			return nil, err
		}
	}
	return []nitro.Value{sumV}, nil
}

type countAccum struct {
	count int
}

func (a *countAccum) String() string    { return "<count>" }
func (a *countAccum) Type() string      { return "count" }
func (a *countAccum) Traits() vm.Traits { return vm.TraitNone }

var errCountUsage = errors.New(
	`invalid usage. Expected count(accum, any) or count(iter) or count(any...)`)

func count(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errSumUsage
	}

	// Form 1: avg(accum|nil, float|int|nil?)
	accum, ok := args[0].(*countAccum)
	if ok || args[0] == nil {
		if len(args) != 1 && len(args) != 2 {
			return nil, errCountUsage
		}
		if accum == nil {
			accum = new(countAccum)
		}
		if len(args) == 1 || args[1] == nil {
			return []nitro.Value{nitro.NewInt(int64(accum.count))}, nil
		}
		accum.count++
		return []nitro.Value{accum}, nil
	}

	if args[0] != nil {
		iter, err := nitro.MakeIterator(m, args[0])
		if err == nil {
			c := 0
			for {
				v, err := m.IterNext(iter, 1)
				if err != nil {
					return nil, err
				}
				if v == nil {
					break
				}
				if v[0] == nil {
					continue
				}
				c++
			}
			return []nitro.Value{nitro.NewInt(int64(c))}, nil
		}
	}

	c := 0
	for _, arg := range args {
		if arg == nil {
			continue
		}
		c++
	}
	return []nitro.Value{nitro.NewInt(int64(c))}, nil
}

var errGroupUsage = errors.New(
	`invalid usage. Expected group(accum, any?)`)

func group(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errGroupUsage
	}

	accum, ok := args[0].(*nitro.Array)
	if !ok && args[0] != nil {
		return nil, errGroupUsage
	}

	if accum == nil {
		accum = nitro.NewArray()
	}
	if len(args) == 1 || args[1] == nil {
		return []nitro.Value{accum}, nil
	}
	accum.Add(args[1])
	return []nitro.Value{accum}, nil
}
