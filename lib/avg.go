package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var errAvgUsage = errors.New(
	"invalid usage. Expected: avg(iter) or avg(accum, int|float)")

type avgAccum struct {
	sum   float64
	count int
}

func (a *avgAccum) String() string { return "<avg>" }
func (a *avgAccum) Type() string   { return "avg" }

func avg(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errAvgUsage
	}

	// Form 1: avg(accum, float|int)
	accum, ok := args[0].(*avgAccum)
	if ok || args[0] == nil {
		if accum == nil {
			accum = new(avgAccum)
		}
		if len(args) != 2 {
			return nil, errAvgUsage
		}
		if args[1] == nil {
			if accum.count <= 0 {
				return nil, fmt.Errorf("accum.count is zero")
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
				"expected arg 2 to be Int|Float, but it was %q",
				nitro.TypeName(v))
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
					"expected arg %d to be Int|Float, but it was %q",
					i, nitro.TypeName(v))
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
		v, ok, err := nitro.Next(m, iter, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
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
				"expected iterator to return only Int|Float, but returned %q",
				nitro.TypeName(v))
		}
	}

	if count == 0 {
		return []nitro.Value{nitro.NewFloat(0)}, nil
	}

	return []nitro.Value{nitro.NewFloat(sum / float64(count))}, nil
}
