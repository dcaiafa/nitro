package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/dcaiafa/nitro"
)

var errDivByZero = errors.New("divide by zero")
var errInvalidOp = errors.New("invalid operation")

type Duration struct {
	dur time.Duration
}

func (d Duration) String() string { return d.dur.String() }
func (d Duration) Type() string   { return "duration" }

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
		} else {
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operand))
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
				return nil, errDivByZero
			}
			return Duration{d.dur % otherDur}, nil
		}

	case nitro.OpMult:
		operandInt, ok := operand.(nitro.Int)
		if !ok {
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operandInt))
		}
		return Duration{d.dur * time.Duration(operandInt.Int64())}, nil

	case nitro.OpDiv:
		switch operand := operand.(type) {
		case Duration:
			if operand.dur == 0 {
				return nil, errDivByZero
			}
			return nitro.NewInt(int64(d.dur / operand.dur)), nil

		case nitro.Int:
			if operand.Int64() == 0 {
				return nil, errDivByZero
			}
			return Duration{d.dur / time.Duration(operand.Int64())}, nil

		default:
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operand))
		}
	}

	return nil, fmt.Errorf("operation not supported by duration")
}

func (d Duration) Duration() time.Duration {
	return d.dur
}

func (d Duration) Index(key nitro.Value) (nitro.Value, error) {
	keyStr, ok := key.(nitro.String)
	if !ok {
		return nil, fmt.Errorf(
			"duration cannot be indexed by %q",
			nitro.TypeName(key))
	}

	switch keyStr.String() {
	case "truncate":
		return nitro.NativeFn(d.truncate), nil
	default:
		return nil, fmt.Errorf(
			"duration does not have method %q",
			keyStr.String())
	}
}

func (d Duration) IndexRef(key nitro.Value) (nitro.ValueRef, error) {
	return nitro.ValueRef{}, fmt.Errorf("duration is read-only")
}

var errDurationTruncateUsage = errors.New(
	`invalid usage. Expected: <duration>.truncate(duration|string)`)

func (d Duration) truncate(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errDurationTruncateUsage
	}

	var mult time.Duration
	if multArg, ok := args[0].(Duration); ok {
		mult = multArg.dur
	} else if multStr, ok := args[0].(nitro.String); ok {
		unit, ok := durationUnit(multStr.String())
		if !ok {
			return nil, errDurationTruncateUsage
		}
		mult = unit
	} else {
		return nil, errDurationTruncateUsage
	}
	truncDur := d.dur.Truncate(mult)
	return []nitro.Value{Duration{truncDur}}, nil
}

var errDurUsage = errors.New(
	`invalid usage. Expected: duration(int, "nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour"`)

func dur(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errDurUsage
	}

	count, ok := args[0].(nitro.Int)
	if !ok {
		return nil, errDurUsage
	}
	unitStr, ok := args[1].(nitro.String)
	if !ok {
		return nil, errDurUsage
	}

	dur := time.Duration(count.Int64())
	unitDur, ok := durationUnit(unitStr.String())
	if !ok {
		return nil, errDurUsage
	}
	dur *= unitDur

	return []nitro.Value{Duration{dur: dur}}, nil
}

func durationUnit(c string) (time.Duration, bool) {
	switch c {
	case "nanosecond":
		return time.Nanosecond, true
	case "microsecond":
		return time.Microsecond, true
	case "millisecond":
		return time.Millisecond, true
	case "second":
		return time.Second, true
	case "minute":
		return time.Minute, true
	case "hour":
		return time.Hour, true
	default:
		return 0, false
	}
}
