package time

import _p0 "github.com/dcaiafa/nitro/internal/export"
import _p1 "github.com/dcaiafa/nitro/internal/stub"
import _p2 "github.com/dcaiafa/nitro/internal/vm"

func _fixed_zone(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.Int:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).Int64()
				_r0, err := fixed_zone0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _format(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) == 1 {
			var _a1 _p2.String = _p2.NewString("2006-01-02T15:04:05Z07:00")
			{
				_ta0 := _a0
				_ta1 := (_a1).String()
				_r0, err := format0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_p2.NewString(_r0)}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := _a0
				_ta1 := (_a1).String()
				_r0, err := format0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_p2.NewString(_r0)}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _from_unix(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.Int:
		if len(args) == 1 {
			var _a1 _p2.Int = _p2.NewInt(0)
			{
				_ta0 := (_a0).Int64()
				_ta1 := (_a1).Int64()
				_r0, err := from_unix0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case _p2.Int:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).Int64()
				_ta1 := (_a1).Int64()
				_r0, err := from_unix0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _in(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		switch _a1 := args[1].(type) {
		case *Location:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := _a0
				_ta1 := _a1
				_r0, err := in0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := _a0
				_ta1 := (_a1).String()
				_r0, err := in1(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _load_location(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := (_a0).String()
			_r0, err := load_location0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _local(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			_r0, err := local0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _now(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) > 0 {
		return nil, _p1.ErrTooManyArgs
	}
	{
		_r0, err := now0(vm)
		if err != nil {
			return nil, err
		}
		return []_p2.Value{_r0}, nil
	}
}
func _parse(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		if len(args) == 1 {
			var _a1 _p2.String = _p2.NewString("2006-01-02T15:04:05Z07:00")
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := parse0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := parse0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _to_map(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			_r0, err := to_map0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _truncate(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Duration:
		switch _a1 := args[1].(type) {
		case Duration:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := _a0
				_ta1 := _a1
				_r0, err := truncate0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _unix(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			_r0, err := unix0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_p2.NewInt(_r0)}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _unix_nano(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			_r0, err := unix_nano0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_p2.NewInt(_r0)}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _utc(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case Time:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			_r0, err := utc0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}

var Exports = _p0.Exports{
	{N: "fixed_zone", T: _p0.Func, F: _fixed_zone},
	{N: "format", T: _p0.Func, F: _format},
	{N: "from_unix", T: _p0.Func, F: _from_unix},
	{N: "in", T: _p0.Func, F: _in},
	{N: "load_location", T: _p0.Func, F: _load_location},
	{N: "local", T: _p0.Func, F: _local},
	{N: "now", T: _p0.Func, F: _now},
	{N: "parse", T: _p0.Func, F: _parse},
	{N: "to_map", T: _p0.Func, F: _to_map},
	{N: "truncate", T: _p0.Func, F: _truncate},
	{N: "unix", T: _p0.Func, F: _unix},
	{N: "unix_nano", T: _p0.Func, F: _unix_nano},
	{N: "utc", T: _p0.Func, F: _utc},
}
