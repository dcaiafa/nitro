package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/dcaiafa/nitro"
)

type Duration struct {
	dur time.Duration
}

func (d Duration) String() string { return d.dur.String() }
func (d Duration) Type() string   { return "duration" }

func (d Duration) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	switch op {
	case nitro.BinEq:
		return nitro.NewBool(d == operand), nil
	case nitro.BinNE:
		return nitro.NewBool(d != operand), nil
	}

	operandDur, ok := operand.(Duration)
	if !ok {
		return nil, fmt.Errorf(
			"invalid operation between duration and %v", nitro.TypeName(operand))
	}

	switch op {
	case nitro.BinAdd:
		return Duration{d.dur + operandDur.dur}, nil
	case nitro.BinSub:
		return Duration{d.dur - operandDur.dur}, nil
	case nitro.BinMult:
		return Duration{d.dur * operandDur.dur}, nil
	case nitro.BinDiv:
		return Duration{d.dur / operandDur.dur}, nil
	case nitro.BinMod:
		return Duration{d.dur % operandDur.dur}, nil
	case nitro.BinLT:
		return nitro.NewBool(d.dur < operandDur.dur), nil
	case nitro.BinLE:
		return nitro.NewBool(d.dur <= operandDur.dur), nil
	case nitro.BinGT:
		return nitro.NewBool(d.dur > operandDur.dur), nil
	case nitro.BinGE:
		return nitro.NewBool(d.dur >= operandDur.dur), nil
	case nitro.BinEq:
		return nitro.NewBool(d.dur == operandDur.dur), nil
	case nitro.BinNE:
		return nitro.NewBool(d.dur != operandDur.dur), nil
	default:
		return nil, fmt.Errorf("operation not supported by duration")
	}
}

func (d Duration) EvalUnaryMinus() (nitro.Value, error) {
	d.dur *= -1
	return d, nil
}

var errDurationUsage = errors.New(
	`invalid usage. Expected: duration(int, "nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour"`)

func dur(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	count, err := getIntArg(args, 0)
	if err != nil {
		return nil, err
	}
	unitStr, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	dur := time.Duration(count)
	unitDur, ok := durationUnit(unitStr)
	if !ok {
		return nil, errDurationUsage
	}
	dur *= unitDur

	return []nitro.Value{Duration{dur: dur}}, nil
}

var errDurationToUsage = errors.New(
	`invalid usage. Expected: durationto(duration, "nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour"`)

func durto(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errDurationToUsage
	}

	dur, ok := args[0].(Duration)
	if !ok {
		return nil, errDurationToUsage
	}
	unitStr, ok := args[1].(nitro.String)
	if !ok {
		return nil, errDurationToUsage
	}
	unitDur, ok := durationUnit(unitStr.String())
	if !ok {
		return nil, errDurationUsage
	}

	count := dur.dur / unitDur
	return []nitro.Value{nitro.NewInt(int64(count))}, nil
}

var errDurationTruncateUsage = errors.New(
	`invalid usage. Expected: durationtruncate(duration, mult: duration|"nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour")`)

func truncdur(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errDurationTruncateUsage
	}

	dur, ok := args[0].(Duration)
	if !ok {
		return nil, errDurationTruncateUsage
	}

	var mult time.Duration
	multArg, ok := args[1].(Duration)
	if ok {
		mult = multArg.dur
	} else {
		multStr, ok := args[1].(nitro.String)
		if !ok {
			return nil, errDurationTruncateUsage
		}
		unit, ok := durationUnit(multStr.String())
		if !ok {
			return nil, errDurationTruncateUsage
		}
		mult = unit
	}

	truncDur := dur.dur.Truncate(mult)

	return []nitro.Value{Duration{truncDur}}, nil
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
