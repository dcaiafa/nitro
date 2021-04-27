package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

func Range(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
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

	i := &rangeIter{
		cur:  start,
		end:  end,
		step: step,
	}

	return []vm.Value{vm.NewIterator(i.Next, 1)}, nil
}

type rangeIter struct {
	cur  int64
	end  int64
	step int64
}

func (i *rangeIter) Next(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error) {
	if (i.step > 0 && i.cur >= i.end) ||
		(i.step < 0 && i.cur <= i.end) ||
		(i.step == 0) {
		return []vm.Value{vm.NewBool(false), vm.NewInt(0)}, nil
	}

	cur := i.cur
	i.cur = i.cur + i.step

	return []vm.Value{vm.NewBool(true), vm.NewInt(cur)}, nil
}
