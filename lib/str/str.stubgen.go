package str

import _p0 "github.com/dcaiafa/nitro/internal/export"
import _p1 "github.com/dcaiafa/nitro/internal/stub"
import _p2 "github.com/dcaiafa/nitro/internal/vm"

func _fields(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
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
			_r0, err := fields0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _find(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := find0(vm, _ta0, _ta1)
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
func _find_last(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := find_last0(vm, _ta0, _ta1)
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
func _has_prefix(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := has_prefix0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_p2.NewBool(_r0)}, nil
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _match(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case *_p2.Regex:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := _a1
				_r0, err := match0(vm, _ta0, _ta1)
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
func _match_all(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case *_p2.Regex:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := _a1
				_r0, err := match_all0(vm, _ta0, _ta1)
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
func _repeat(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
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
				_r0, err := repeat0(vm, _ta0, _ta1)
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
func _replace(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 3 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case *_p2.Regex:
			switch _a2 := args[2].(type) {
			case _p2.Callable:
				if len(args) > 3 {
					return nil, _p1.ErrTooManyArgs
				}
				{
					_ta0 := (_a0).String()
					_ta1 := _a1
					_ta2 := _a2
					_r0, err := replace2(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_p2.NewString(_r0)}, nil
				}
			case _p2.String:
				if len(args) > 3 {
					return nil, _p1.ErrTooManyArgs
				}
				{
					_ta0 := (_a0).String()
					_ta1 := _a1
					_ta2 := (_a2).String()
					_r0, err := replace1(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_p2.NewString(_r0)}, nil
				}
			default:
				return nil, _p1.InvalidArg(args, 2)
			}
		case _p2.String:
			switch _a2 := args[2].(type) {
			case _p2.String:
				if len(args) == 3 {
					var _a3 _p2.Int = _p2.NewInt(-1)
					{
						_ta0 := (_a0).String()
						_ta1 := (_a1).String()
						_ta2 := (_a2).String()
						_ta3 := (_a3).Int64()
						_r0, err := replace0(vm, _ta0, _ta1, _ta2, _ta3)
						if err != nil {
							return nil, err
						}
						return []_p2.Value{_p2.NewString(_r0)}, nil
					}
				}
				switch _a3 := args[3].(type) {
				case _p2.Int:
					if len(args) > 4 {
						return nil, _p1.ErrTooManyArgs
					}
					{
						_ta0 := (_a0).String()
						_ta1 := (_a1).String()
						_ta2 := (_a2).String()
						_ta3 := (_a3).Int64()
						_r0, err := replace0(vm, _ta0, _ta1, _ta2, _ta3)
						if err != nil {
							return nil, err
						}
						return []_p2.Value{_p2.NewString(_r0)}, nil
					}
				default:
					return nil, _p1.InvalidArg(args, 3)
				}
			default:
				return nil, _p1.InvalidArg(args, 2)
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _split(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case *_p2.Regex:
			if len(args) == 2 {
				var _a2 _p2.Int = _p2.NewInt(-1)
				{
					_ta0 := (_a0).String()
					_ta1 := _a1
					_ta2 := (_a2).Int64()
					_r0, err := split1(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_r0}, nil
				}
			}
			switch _a2 := args[2].(type) {
			case _p2.Int:
				if len(args) > 3 {
					return nil, _p1.ErrTooManyArgs
				}
				{
					_ta0 := (_a0).String()
					_ta1 := _a1
					_ta2 := (_a2).Int64()
					_r0, err := split1(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_r0}, nil
				}
			default:
				return nil, _p1.InvalidArg(args, 2)
			}
		case _p2.String:
			if len(args) == 2 {
				var _a2 _p2.Int = _p2.NewInt(-1)
				{
					_ta0 := (_a0).String()
					_ta1 := (_a1).String()
					_ta2 := (_a2).Int64()
					_r0, err := split0(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_r0}, nil
				}
			}
			switch _a2 := args[2].(type) {
			case _p2.Int:
				if len(args) > 3 {
					return nil, _p1.ErrTooManyArgs
				}
				{
					_ta0 := (_a0).String()
					_ta1 := (_a1).String()
					_ta2 := (_a2).Int64()
					_r0, err := split0(vm, _ta0, _ta1, _ta2)
					if err != nil {
						return nil, err
					}
					return []_p2.Value{_r0}, nil
				}
			default:
				return nil, _p1.InvalidArg(args, 2)
			}
		default:
			return nil, _p1.InvalidArg(args, 1)
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _to_lower(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
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
			_r0, err := to_lower0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_p2.NewString(_r0)}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _to_upper(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
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
			_r0, err := to_upper0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_p2.NewString(_r0)}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _trim_prefix(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := trim_prefix0(vm, _ta0, _ta1)
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
func _trim_space(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
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
			_r0, err := trim_space0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_p2.NewString(_r0)}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _trim_suffix(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		switch _a1 := args[1].(type) {
		case _p2.String:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				_ta1 := (_a1).String()
				_r0, err := trim_suffix0(vm, _ta0, _ta1)
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

var Exports = _p0.Exports{
	{N: "fields", T: _p0.Func, F: _fields},
	{N: "find", T: _p0.Func, F: _find},
	{N: "find_last", T: _p0.Func, F: _find_last},
	{N: "has_prefix", T: _p0.Func, F: _has_prefix},
	{N: "match", T: _p0.Func, F: _match},
	{N: "match_all", T: _p0.Func, F: _match_all},
	{N: "repeat", T: _p0.Func, F: _repeat},
	{N: "replace", T: _p0.Func, F: _replace},
	{N: "split", T: _p0.Func, F: _split},
	{N: "to_lower", T: _p0.Func, F: _to_lower},
	{N: "to_upper", T: _p0.Func, F: _to_upper},
	{N: "trim_prefix", T: _p0.Func, F: _trim_prefix},
	{N: "trim_space", T: _p0.Func, F: _trim_space},
	{N: "trim_suffix", T: _p0.Func, F: _trim_suffix},
}
