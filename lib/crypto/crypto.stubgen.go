package crypto

import (
	_p0 "github.com/dcaiafa/nitro"
	_p1 "github.com/dcaiafa/nitro/internal/export"
	_p2 "github.com/dcaiafa/nitro/internal/stub"

	_p3 "github.com/dcaiafa/nitro/internal/vm"
)

func _generate_secret(vm *_p3.VM, args []_p3.Value, nret int) ([]_p3.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p2.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p3.Int:
		if len(args) == 1 {
			var _a1 _p3.String = _p3.NewString("")
			{
				_ta0 := (_a0).Int64()
				_ta1 := (_a1).String()
				_r0, err := generate_secret0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p3.Value{_p3.NewString(_r0)}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case _p3.String:
			if len(args) > 2 {
				return nil, _p2.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).Int64()
				_ta1 := (_a1).String()
				_r0, err := generate_secret0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p3.Value{_p3.NewString(_r0)}, nil
			}
		default:
			return nil, _p2.InvalidArg(args, 1)
		}
	default:
		return nil, _p2.InvalidArg(args, 0)
	}
}

var Exports = _p1.Exports{
	{N: "generate_secret", T: _p1.Func, F: _generate_secret},
	{N: "ALPHANUM", T: _p1.Value, V: _p0.NewString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")},
}
