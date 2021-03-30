package std

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func Range(ctx context.Context, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	var err error
	var start int64 = 0
	var step int64 = 0
	var end int64

	if len(args) == 0 {
		return nil, errNotEnoughArgs
	} else if len(args) > 3 {
		return nil, errTooManyArgs
	}

	if len(args) == 1 {
		end, err = getIntArg(args, 0)
		if err != nil {
			return nil, err
		}
	} else if len(args) >= 2 {
		start, err = getIntArg(args, 0)
		if err != nil {
			return nil, err
		}
		end, err = getIntArg(args, 1)
		if err != nil {
			return nil, err
		}
		if len(args) == 3 {
			step, err = getIntArg(args, 2)
			if err != nil {
				return nil, err
			}
			if step == 0 {
				return nil, fmt.Errorf("step argument cannot be zero")
			}
		}
	}

	if step == 0 {
		if start < end {
			step = 1
		} else if start > end {
			step = -1
		}
	}

	if start < end {
		if step < 0 {
			return nil, fmt.Errorf("when start < end, step must be positive")
		}
	} else if start > end && step > 0 {
		if step < 0 {
			return nil, fmt.Errorf("when start > end, step must be negative")
		}
	}

	var (
		currentValue runtime.Value = runtime.NewInt(start)
		endValue     runtime.Value = runtime.NewInt(end)
		stepValue    runtime.Value = runtime.NewInt(step)
	)

	c := runtime.NewClosure(rangeIter, []runtime.ValueRef{
		runtime.NewValueRef(&currentValue),
		runtime.NewValueRef(&endValue),
		runtime.NewValueRef(&stepValue),
	})

	return []runtime.Value{c}, nil
}

func rangeIter(ctx context.Context, caps []runtime.ValueRef, args []runtime.Value, expRetN int) ([]runtime.Value, error) {
	var (
		cur  = ((*caps[0].Ref).(runtime.Int)).Int64()
		end  = ((*caps[1].Ref).(runtime.Int)).Int64()
		step = ((*caps[2].Ref).(runtime.Int)).Int64()
	)

	if (step > 0 && cur >= end) ||
		(step < 0 && cur <= end) ||
		(step == 0) {
		return []runtime.Value{runtime.NewBool(false), runtime.NewInt(0)}, nil
	}

	*caps[0].Refo() = runtime.NewInt(cur + step)

	return []runtime.Value{runtime.NewBool(true), runtime.NewInt(cur)}, nil
}
