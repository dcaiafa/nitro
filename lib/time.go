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

var _ /* implements */ nitro.Indexable = Time{}

func NewTime(t time.Time) Time {
	return Time{time: t}
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

func (t Time) Index(key nitro.Value) (nitro.Value, error) {
	keyStr, ok := key.(nitro.String)
	if !ok {
		return nil, fmt.Errorf(
			"time cannot be indexed by %q",
			nitro.TypeName(key))
	}

	switch keyStr.String() {
	case "utc":
		return nitro.NativeFn(t.utc), nil
	case "local":
		return nitro.NativeFn(t.local), nil
	case "format":
		return nitro.NativeFn(t.format), nil
	case "unix":
		return nitro.NativeFn(t.unix), nil
	case "unixnano":
		return nitro.NativeFn(t.unixnano), nil
	case "year":
		return nitro.NewInt(int64(t.time.Year())), nil
	case "month":
		return nitro.NewInt(int64(t.time.Month())), nil
	case "day":
		return nitro.NewInt(int64(t.time.Day())), nil
	case "hour":
		return nitro.NewInt(int64(t.time.Hour())), nil
	case "minute":
		return nitro.NewInt(int64(t.time.Minute())), nil
	case "second":
		return nitro.NewInt(int64(t.time.Second())), nil
	case "nanosecond":
		return nitro.NewInt(int64(t.time.Nanosecond())), nil
	default:
		return nil, fmt.Errorf(
			"time does not have member %q",
			keyStr.String())
	}
}

func (t Time) IndexRef(key nitro.Value) (nitro.ValueRef, error) {
	return nitro.ValueRef{}, fmt.Errorf("time is read-only")
}

func (t Time) utc(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("utc takes no arguments")
	}
	t.time = t.time.UTC()
	return []nitro.Value{t}, nil
}

func (t Time) local(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("local takes no arguments")
	}
	t.time = t.time.Local()
	return []nitro.Value{t}, nil
}

var errTimeFormatUsage = errors.New(
	"invalid usage. Expected <time>.format(layout: string?)")

func (t Time) format(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 && len(args) != 1 {
		return nil, errTimeFormatUsage
	}
	layout := time.RFC3339
	if len(args) == 1 {
		layoutArg, ok := args[0].(nitro.String)
		if !ok {
			return nil, errTimeFormatUsage
		}
		layout = convertTimeLayout(layoutArg.String())
	}
	res := t.time.Format(layout)
	return []nitro.Value{nitro.NewString(res)}, nil
}

func (t Time) unix(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errTakesNoArgs
	}
	return []nitro.Value{nitro.NewInt(t.time.Unix())}, nil
}

func (t Time) unixnano(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errTakesNoArgs
	}
	return []nitro.Value{nitro.NewInt(t.time.UnixNano())}, nil
}

func now(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errTakesNoArgs
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
