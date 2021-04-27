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

func (d Duration) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	switch op {
	case nitro.BinEq:
		return nitro.NewBool(d == operand), nil
	case nitro.BinNE:
		return nitro.NewBool(d != operand), nil
	}

	switch op {
	case nitro.BinAdd, nitro.BinSub, nitro.BinLT, nitro.BinLE,
		nitro.BinGT, nitro.BinGE, nitro.BinMod:

		operandDur, ok := operand.(Duration)
		if !ok {
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operand))
		}

		switch op {
		case nitro.BinAdd:
			return Duration{d.dur + operandDur.dur}, nil
		case nitro.BinSub:
			return Duration{d.dur - operandDur.dur}, nil
		case nitro.BinLT:
			return nitro.NewBool(d.dur < operandDur.dur), nil
		case nitro.BinLE:
			return nitro.NewBool(d.dur <= operandDur.dur), nil
		case nitro.BinGT:
			return nitro.NewBool(d.dur > operandDur.dur), nil
		case nitro.BinGE:
			return nitro.NewBool(d.dur >= operandDur.dur), nil
		case nitro.BinMod:
			if operandDur.dur == 0 {
				return nil, errDivByZero
			}
			return Duration{d.dur % operandDur.dur}, nil
		}

	case nitro.BinMult:
		operandInt, ok := operand.(nitro.Int)
		if !ok {
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operandInt))
		}
		return Duration{d.dur * time.Duration(operandInt.Int64())}, nil

	case nitro.BinDiv:
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

func (d Duration) EvalUnaryMinus() (nitro.Value, error) {
	d.dur *= -1
	return d, nil
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

var errTruncDurUsage = errors.New(
	`invalid usage. Expected: truncdur(duration, mult: duration|"nanosecond"|"microsecond"|"millisecond"|"second"|"minute"|"hour")`)

func truncdur(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errTruncDurUsage
	}

	dur, ok := args[0].(Duration)
	if !ok {
		return nil, errTruncDurUsage
	}

	var mult time.Duration
	multArg, ok := args[1].(Duration)
	if ok {
		mult = multArg.dur
	} else {
		multStr, ok := args[1].(nitro.String)
		if !ok {
			return nil, errTruncDurUsage
		}
		unit, ok := durationUnit(multStr.String())
		if !ok {
			return nil, errTruncDurUsage
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
