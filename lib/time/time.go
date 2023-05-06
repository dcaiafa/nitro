package time

import (
	"fmt"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/internal/vm"
)

//go:generate stubgen time.stubgen

func init() {
	Exports = append(Exports,
		export.Export{N: "HOUR", T: export.Value, V: NewDuration(time.Hour)},
		export.Export{N: "MICROSECOND", T: export.Value, V: NewDuration(time.Microsecond)},
		export.Export{N: "MILLISECOND", T: export.Value, V: NewDuration(time.Millisecond)},
		export.Export{N: "MINUTE", T: export.Value, V: NewDuration(time.Minute)},
		export.Export{N: "NANOSECOND", T: export.Value, V: NewDuration(time.Nanosecond)},
		export.Export{N: "SECOND", T: export.Value, V: NewDuration(time.Second)},
	)
}

type Time struct {
	time time.Time
}

var _ /* implements */ nitro.Value = Time{}

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

func utc0(m *vm.VM, t Time) (Time, error) {
	t.time = t.time.UTC()
	return t, nil
}

func local0(m *vm.VM, t Time) (Time, error) {
	t.time = t.time.Local()
	return t, nil
}

func in0(m *vm.VM, t Time, loc *Location) (Time, error) {
	return NewTime(t.Time().In(loc.Location)), nil
}

func in1(m *vm.VM, t Time, loc string) (Time, error) {
	cache := getLocationCache(m)
	locObj, err := cache.GetLocation(loc)
	if err != nil {
		return Time{}, err
	}
	return in0(m, t, locObj)
}

func format0(m *vm.VM, t Time, layout string) (string, error) {
	return t.time.Format(layout), nil
}

func unix0(m *vm.VM, t Time) (int64, error) {
	return t.time.Unix(), nil
}

func unix_nano0(m *vm.VM, t Time) (int64, error) {
	return t.time.UnixNano(), nil
}

func now0(m *vm.VM) (Time, error) {
	return NewTime(time.Now()), nil
}

func parse0(m *vm.VM, v string, layout string) (Time, error) {
	t, err := time.Parse(layout, v)
	if err != nil {
		return Time{}, err
	}
	return NewTime(t), nil
}

func from_unix0(m *vm.VM, sec, nano int64) (Time, error) {
	t := time.Unix(sec, nano)
	return NewTime(t), nil
}

func to_map0(vm *vm.VM, t Time) (*vm.Object, error) {
	m := nitro.NewObject()
	m.Put(nitro.NewString("year"), nitro.NewInt(int64(t.time.Year())))
	m.Put(nitro.NewString("month"), nitro.NewInt(int64(t.time.Month())))
	m.Put(nitro.NewString("day"), nitro.NewInt(int64(t.time.Day())))
	m.Put(nitro.NewString("hour"), nitro.NewInt(int64(t.time.Hour())))
	m.Put(nitro.NewString("minute"), nitro.NewInt(int64(t.time.Minute())))
	m.Put(nitro.NewString("second"), nitro.NewInt(int64(t.time.Second())))
	m.Put(nitro.NewString("nanosecond"), nitro.NewInt(int64(t.time.Nanosecond())))
	return m, nil
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

func truncate0(vm *nitro.VM, d Duration, m Duration) (Duration, error) {
	v := d.dur.Truncate(m.dur)
	return NewDuration(v), nil
}

type Location struct {
	Location *time.Location
}

func (l *Location) String() string    { return l.Location.String() }
func (l *Location) Type() string      { return "location" }
func (l *Location) Traits() vm.Traits { return vm.TraitNone }

func fixed_zone0(vm *vm.VM, name string, offset int64) (*Location, error) {
	return &Location{
		Location: time.FixedZone(name, int(offset)),
	}, nil
}

type locationCache struct {
	locations map[string]*Location
}

const locationCacheUserDataKey = "locationCache"

func getLocationCache(m *vm.VM) *locationCache {
	cache, ok := m.GetUserData(locationCacheUserDataKey).(*locationCache)
	if !ok {
		cache = &locationCache{
			locations: make(map[string]*Location),
		}
		m.SetUserData(locationCacheUserDataKey, cache)
	}
	return cache
}

func (c *locationCache) GetLocation(name string) (*Location, error) {
	loc := c.locations[name]
	if loc != nil {
		return loc, nil
	}

	timeLocation, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	location := &Location{
		Location: timeLocation,
	}

	c.locations[name] = location

	return location, nil
}

func load_location0(vm *vm.VM, name string) (*Location, error) {
	loc, err := getLocationCache(vm).GetLocation(name)
	if err != nil {
		return nil, err
	}
	return loc, nil
}
