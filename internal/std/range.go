package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

func Range(m *vm.VM, caps []vm.ValueRef, args []vm.Value, nRet int) ([]vm.Value, error) {
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
		currentValue vm.Value = vm.NewInt(start)
		endValue     vm.Value = vm.NewInt(end)
		stepValue    vm.Value = vm.NewInt(step)
	)

	c := vm.NewIterator(
		rangeIter,
		[]vm.ValueRef{
			vm.NewValueRef(&currentValue),
			vm.NewValueRef(&endValue),
			vm.NewValueRef(&stepValue),
		},
		1)

	return []vm.Value{c}, nil
}

func rangeIter(m *vm.VM, caps []vm.ValueRef, args []vm.Value, nRet int) ([]vm.Value, error) {
	var (
		cur  = ((*caps[0].Ref).(vm.Int)).Int64()
		end  = ((*caps[1].Ref).(vm.Int)).Int64()
		step = ((*caps[2].Ref).(vm.Int)).Int64()
	)

	if (step > 0 && cur >= end) ||
		(step < 0 && cur <= end) ||
		(step == 0) {
		return []vm.Value{vm.NewBool(false), vm.NewInt(0)}, nil
	}

	*caps[0].Refo() = vm.NewInt(cur + step)

	return []vm.Value{vm.NewBool(true), vm.NewInt(cur)}, nil
}
