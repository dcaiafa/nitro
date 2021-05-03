package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/dcaiafa/nitro"
)

type Time struct {
	time time.Time
}

func (t Time) String() string { return t.time.String() }
func (t Time) Type() string   { return "time" }

func (t Time) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	switch op {
	case nitro.OpEq, nitro.OpSub, nitro.OpLT, nitro.OpLE, nitro.OpGT, nitro.OpGE:
		operandTime, ok := operand.(Time)
		if !ok {
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
		return nil, fmt.Errorf("operation is not supported by time")
	}
}

var errNowUsage = errors.New("invalid usage. Expected now()")

func now(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errNowUsage
	}
	return []nitro.Value{Time{time.Now()}}, nil
}

func parsetime(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	timeStr, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	var layout string
	if len(args) >= 2 {
		layout, err = getStringArg(args, 1)
		if err != nil {
			return nil, err
		}
		layout = convertTimeLayout(layout)
	} else {
		layout = time.RFC3339
	}
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{Time{time: t}}, nil
}

var errFormatTimeUsage = errors.New(
	"invalid usage. Expected formattime(time, layout: string?)")

func formattime(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errFormatTimeUsage
	}
	t, ok := args[0].(Time)
	if !ok {
		return nil, errFormatTimeUsage
	}
	layout := time.RFC3339
	if len(args) == 2 {
		layoutArg, ok := args[1].(nitro.String)
		if !ok {
			return nil, errFormatTimeUsage
		}
		layout = convertTimeLayout(layoutArg.String())
	}

	res := t.time.Format(layout)

	return []nitro.Value{nitro.NewString(res)}, nil
}

func timetounix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	timeArg, ok := args[0].(Time)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument 1 to be time, but it is %v",
			nitro.TypeName(args[0]))
	}

	return []nitro.Value{nitro.NewInt(timeArg.time.Unix())}, nil
}

func timetounixnano(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	timeArg, ok := args[0].(Time)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument 1 to be time, but it is %v",
			nitro.TypeName(args[0]))
	}

	return []nitro.Value{nitro.NewInt(timeArg.time.UnixNano())}, nil
}

func timefromunix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	var sec, nano int64
	sec, err = getIntArg(args, 0)
	if err != nil {
		return nil, err
	}
	if len(args) >= 2 {
		nano, err = getIntArg(args, 1)
		if err != nil {
			return nil, err
		}
	}

	t := time.Unix(sec, nano)

	return []nitro.Value{Time{time: t}}, nil
}

func convertTimeLayout(timeFmt string) string {
	switch timeFmt {
	case "ansic":
		return time.ANSIC
	case "unixdate":
		return time.UnixDate
	case "rubydate":
		return time.RubyDate
	case "rfc822":
		return time.RFC822
	case "rfc822z":
		return time.RFC822Z
	case "rfc850":
		return time.RFC850
	case "rfc1123":
		return time.RFC1123
	case "rfc1123z":
		return time.RFC1123Z
	case "rfc3339":
		return time.RFC3339
	case "rfc3339nano":
		return time.RFC3339Nano
	case "kitchen":
		return time.Kitchen
	case "stamp":
		return time.Stamp
	case "stampmilli":
		return time.StampMilli
	case "stampmicro":
		return time.StampMicro
	case "stampnano":
		return time.StampNano
	default:
		return timeFmt
	}
}
