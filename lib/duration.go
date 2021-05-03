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

		operandDur, ok := operand.(Duration)
		if !ok {
			return nil, fmt.Errorf(
				"%w between duration and %v", errInvalidOp, nitro.TypeName(operand))
		}

		switch op {
		case nitro.OpAdd:
			return Duration{d.dur + operandDur.dur}, nil
		case nitro.OpSub:
			return Duration{d.dur - operandDur.dur}, nil
		case nitro.OpLT:
			return nitro.NewBool(d.dur < operandDur.dur), nil
		case nitro.OpLE:
			return nitro.NewBool(d.dur <= operandDur.dur), nil
		case nitro.OpGT:
			return nitro.NewBool(d.dur > operandDur.dur), nil
		case nitro.OpGE:
			return nitro.NewBool(d.dur >= operandDur.dur), nil
		case nitro.OpEq:
			return nitro.NewBool(d == operand), nil
		case nitro.OpMod:
			if operandDur.dur == 0 {
				return nil, errDivByZero
			}
			return Duration{d.dur % operandDur.dur}, nil
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
