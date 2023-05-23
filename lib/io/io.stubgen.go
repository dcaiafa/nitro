package io

import _p0 "github.com/dcaiafa/nitro/internal/export"
import _p1 "github.com/dcaiafa/nitro/internal/stub"
import _p2 "github.com/dcaiafa/nitro/internal/vm"

func _err(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) == 0 {
		var _a0 _p2.Reader = nil
		{
			_ta0 := _p1.MustMakeReader(vm, _a0)
			_r0, err := err0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	}
	switch _a0 := args[0].(type) {
	case _p2.Reader, _p2.Readable:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _p1.MustMakeReader(vm, _a0)
			_r0, err := err0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _in(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) > 0 {
		return nil, _p1.ErrTooManyArgs
	}
	{
		_r0, err := in0(vm)
		if err != nil {
			return nil, err
		}
		return []_p2.Value{_r0}, nil
	}
}
func _out(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) == 0 {
		var _a0 _p2.Reader = nil
		{
			_ta0 := _p1.MustMakeReader(vm, _a0)
			_r0, err := out0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	}
	switch _a0 := args[0].(type) {
	case _p2.Reader, _p2.Readable:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _p1.MustMakeReader(vm, _a0)
			_r0, err := out0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _pop_out(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) > 0 {
		return nil, _p1.ErrTooManyArgs
	}
	{
		_r0, err := pop_out0(vm)
		if err != nil {
			return nil, err
		}
		return []_p2.Value{_r0}, nil
	}
}
func _push_out(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.Writer:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _a0
			err := push_out0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}

var Exports = _p0.Exports{
	{N: "err", T: _p0.Func, F: _err},
	{N: "in", T: _p0.Func, F: _in},
	{N: "out", T: _p0.Func, F: _out},
	{N: "pop_out", T: _p0.Func, F: _pop_out},
	{N: "push_out", T: _p0.Func, F: _push_out},
}
