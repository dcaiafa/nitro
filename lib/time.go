package lib

import (
	"fmt"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Time struct {
	time time.Time
}

var _ /* implements */ nitro.Indexable = Time{}

func NewTime(t time.Time) Time {
	return Time{time: t}
}

func (t Time) Time() time.Time { return t.time }

func (t Time) String() string    { return t.time.String() }
func (t Time) Type() string      { return "time" }
func (t Time) Traits() vm.Traits { return vm.TraitEq }

func (t Time) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	switch op {
	case nitro.OpEq, nitro.OpSub, nitro.OpLT, nitro.OpLE, nitro.OpGT, nitro.OpGE:
		operandTime, ok := operand.(Time)
		if !ok {
			if op == nitro.OpEq {
				return nitro.NewBool(false), nil
			}
			return nil, fmt.Errorf(
				"invalid operation between time and %v",
				nitro.TypeName(operand))
		}

		switch op {
		case nitro.OpEq:
			return nitro.NewBool(t.time.Equal(operandTime.time)), nil
		case nitro.OpSub:
			return Duration{t.time.Sub(operandTime.time)}, nil
		case nitro.OpLT:
			return nitro.NewBool(t.time.Before(operandTime.time)), nil
		case nitro.OpLE:
			return nitro.NewBool(t.time.Before(operandTime.time) ||
				t.time.Equal(operandTime.time)), nil
		case nitro.OpGT:
			return nitro.NewBool(t.time.After(operandTime.time)), nil
		case nitro.OpGE:
			return nitro.NewBool(t.time.After(operandTime.time) ||
				t.time.Equal(operandTime.time)), nil
		default:
			panic("unreachable")
		}

	case nitro.OpAdd:
		operandDur, ok := operand.(Duration)
		if !ok {
			return nil, fmt.Errorf(
				"invalid operation between time and %v",
				nitro.TypeName(operand))
		}
		return Time{t.time.Add(operandDur.dur)}, nil

	default:
		return nil, vm.ErrOperationNotSupported
	}
}

func (t Time) Index(key nitro.Value) (nitro.Value, error) {
	return nil, vm.ErrOperationNotSupported
}

func (t Time) IndexRef(key nitro.Value) (nitro.ValueRef, error) {
	return nitro.ValueRef{}, fmt.Errorf("time is read-only")
}

func timeUTC(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}
	t.time = t.time.UTC()
	return []nitro.Value{t}, nil
}

func timeLocal(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}
	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}
	t.time = t.time.Local()
	return []nitro.Value{t}, nil
}

func timeFormat(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}
	layout := time.RFC3339
	if len(args) == 2 {
		layout, err = getStringArg(args, 1)
		if err != nil {
			return nil, err
		}
	}
	res := t.time.Format(layout)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func timeUnix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errInvalidNumberOfArgs
	}
	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewInt(t.time.Unix())}, nil
}

func timeUnixNano(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errInvalidNumberOfArgs
	}
	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewInt(t.time.UnixNano())}, nil
}

func timeNow(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 0 {
		return nil, errTooManyArgs
	}
	return []nitro.Value{Time{time.Now()}}, nil
}

func timeParse(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	timeStr, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	var layout string
	if len(args) == 2 {
		layout, err = getStringArg(args, 1)
		if err != nil {
			return nil, err
		}
	} else {
		layout = time.RFC3339
	}
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{Time{time: t}}, nil
}

func timeFromUnix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}
	var err error
	var sec, nano int64
	sec, err = getIntArg(args, 0)
	if err != nil {
		return nil, err
	}
	if len(args) == 2 {
		nano, err = getIntArg(args, 1)
		if err != nil {
			return nil, err
		}
	}
	t := time.Unix(sec, nano)
	return []nitro.Value{Time{time: t}}, nil
}

func timeToMap(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errTooManyArgs
	}

	t, err := getTimeArg(args, 0)
	if err != nil {
		return nil, err
	}

	m := nitro.NewObject()

	m.Put(nitro.NewString("year"), nitro.NewInt(int64(t.time.Year())))
	m.Put(nitro.NewString("month"), nitro.NewInt(int64(t.time.Month())))
	m.Put(nitro.NewString("day"), nitro.NewInt(int64(t.time.Day())))
	m.Put(nitro.NewString("hour"), nitro.NewInt(int64(t.time.Hour())))
	m.Put(nitro.NewString("minute"), nitro.NewInt(int64(t.time.Minute())))
	m.Put(nitro.NewString("second"), nitro.NewInt(int64(t.time.Second())))
	m.Put(nitro.NewString("nanosecond"), nitro.NewInt(int64(t.time.Nanosecond())))

	return []nitro.Value{m}, nil
}

type Duration struct {
	dur time.Duration
}

func NewDuration(dur time.Duration) Duration {
	return Duration{dur: dur}
}

func (d Duration) String() string    { return d.dur.String() }
func (d Duration) Type() string      { return "duration" }
func (d Duration) Traits() vm.Traits { return vm.TraitEq }

func (d Duration) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	if op == nitro.OpUMinus {
		return Duration{d.dur * -1}, nil
	}

	switch op {
	case nitro.OpAdd, nitro.OpSub, nitro.OpLT, nitro.OpLE,
		nitro.OpGT, nitro.OpGE, nitro.OpEq, nitro.OpMod:

		otherDur := time.Duration(0)
		operandDur, ok := operand.(Duration)
		if ok {
			otherDur = operandDur.dur
		} else if operandInt, ok := operand.(nitro.Int); ok && operandInt.Int64() == 0 {
			// Zero is a special case. It is useful to express `dur < 0` without
			// having to create a duration value for the right side.
			otherDur = 0
		} else if op == nitro.OpEq {
			return nitro.NewBool(false), nil
		} else {
			return nil, vm.ErrOperationNotSupported
		}

		switch op {
		case nitro.OpAdd:
			return Duration{d.dur + otherDur}, nil
		case nitro.OpSub:
			return Duration{d.dur - otherDur}, nil
		case nitro.OpLT:
			return nitro.NewBool(d.dur < otherDur), nil
		case nitro.OpLE:
			return nitro.NewBool(d.dur <= otherDur), nil
		case nitro.OpGT:
			return nitro.NewBool(d.dur > otherDur), nil
		case nitro.OpGE:
			return nitro.NewBool(d.dur >= otherDur), nil
		case nitro.OpEq:
			return nitro.NewBool(d == operand), nil
		case nitro.OpMod:
			if otherDur == 0 {
				return nil, vm.ErrDivByZero
			}
			return Duration{d.dur % otherDur}, nil
		}

	case nitro.OpMult:
		operandInt, ok := operand.(nitro.Int)
		if !ok {
			return nil, vm.ErrOperationNotSupported
		}
		return Duration{d.dur * time.Duration(operandInt.Int64())}, nil

	case nitro.OpDiv:
		switch operand := operand.(type) {
		case Duration:
			if operand.dur == 0 {
				return nil, vm.ErrDivByZero
			}
			return nitro.NewInt(int64(d.dur / operand.dur)), nil

		case nitro.Int:
			if operand.Int64() == 0 {
				return nil, vm.ErrDivByZero
			}
			return Duration{d.dur / time.Duration(operand.Int64())}, nil

		default:
			return nil, vm.ErrOperationNotSupported
		}
	}

	return nil, vm.ErrOperationNotSupported
}

func (d Duration) FallbackEvalOp(op nitro.Op, left nitro.Value) (nitro.Value, error) {
	if op == nitro.OpMult {
		if left, ok := left.(nitro.Int); ok {
			return NewDuration(time.Duration(left.Int64()) * d.dur), nil
		}
	}
	return nil, vm.ErrOperationNotSupported
}

func (d Duration) Duration() time.Duration {
	return d.dur
}

func timeTruncate(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	dur, err := getDurationArg(args, 0)
	if err != nil {
		return nil, err
	}

	mult, err := getDurationArg(args, 1)
	if err != nil {
		return nil, err
	}

	truncDur := dur.dur.Truncate(mult.dur)
	return []nitro.Value{NewDuration(truncDur)}, nil
}
