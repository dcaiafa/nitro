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

func duration(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	count, err := getIntArg(args, 0)
	if err != nil {
		return nil, err
	}
	unit, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	dur := time.Duration(count)
	switch unit {
	case "nanosecond":
		dur *= time.Nanosecond
	case "microsecond":
		dur *= time.Microsecond
	case "millisecond":
		dur *= time.Millisecond
	case "second":
		dur *= time.Second
	case "minute":
		dur *= time.Minute
	case "hour":
		dur *= time.Hour
	default:
		return nil, errDurationUsage
	}

	return []nitro.Value{Duration{dur: dur}}, nil
}

var errDurationToUsage = errors.New(
	`invalid usage. Expected: durationto(duration, "nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour"`)

func durationto(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errDurationToUsage
	}

	dur, ok := args[0].(Duration)
	if !ok {
		return nil, errDurationToUsage
	}
	unit, ok := args[1].(nitro.String)
	if !ok {
		return nil, errDurationToUsage
	}

	count := dur.dur
	switch unit.String() {
	case "nanosecond":
		count /= time.Nanosecond
	case "microsecond":
		count /= time.Microsecond
	case "millisecond":
		count /= time.Millisecond
	case "second":
		count /= time.Second
	case "minute":
		count /= time.Minute
	case "hour":
		count /= time.Hour
	default:
		return nil, errDurationUsage
	}

	return []nitro.Value{nitro.NewInt(int64(count))}, nil
}
