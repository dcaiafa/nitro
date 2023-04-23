package json

import _p0 "github.com/dcaiafa/nitro/internal/export"
import _p1 "github.com/dcaiafa/nitro/internal/stub"
import _p2 "github.com/dcaiafa/nitro/internal/vm"

type EncodeOptions struct {
	Indent string
	Prefix string
}

func (m *EncodeOptions) FromMap(v *_p2.Object) error {
	var err error
	_ = err
	v.ForEach(func(k, v _p2.Value) bool {
		n, ok := k.(_p2.String)
		if !ok {
			err = _p1.ErrMapKeyMustBeStr
			return false
		}
		switch n.String() {
		case "indent":
			cv, ok := v.(_p2.String)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).String()
			m.Indent = tv
		case "prefix":
			cv, ok := v.(_p2.String)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).String()
			m.Prefix = tv
		default:
			err = _p1.StructDoesNotHaveField("EncodeOptions", n.String())
			return false
		}
		return true
	})
	if err != nil {
		return err
	}
	return nil
}
func _decode(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.Reader, _p2.Readable:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := _p1.MustMakeReader(vm, _a0)
			_r0, err := decode0(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	case _p2.String:
		if len(args) > 1 {
			return nil, _p1.ErrTooManyArgs
		}
		{
			_ta0 := (_a0).String()
			_r0, err := decode1(vm, _ta0)
			if err != nil {
				return nil, err
			}
			return []_p2.Value{_r0}, nil
		}
	default:
		return nil, _p1.InvalidArg(args, 0)
	}
}
func _encode(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 1 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.Value:
		if len(args) == 1 {
			var _a1 *_p2.Object = nil
			{
				_ta0 := _a0
				var _ta1 *EncodeOptions
				if _a1 != nil {
					_ta1 = new(EncodeOptions)
					err = _ta1.FromMap(_a1)
					if err != nil {
						return nil, _p1.InvalidArgErr(args, 1, err)
					}
				}
				_r0, err := encode0(vm, _ta0, _ta1)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_p2.NewString(_r0)}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case *_p2.Object:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := _a0
				var _ta1 *EncodeOptions
				if _a1 != nil {
					_ta1 = new(EncodeOptions)
					err = _ta1.FromMap(_a1)
					if err != nil {
						return nil, _p1.InvalidArgErr(args, 2, err)
					}
				}
				_r0, err := encode0(vm, _ta0, _ta1)
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
	{N: "decode", T: _p0.Func, F: _decode},
	{N: "encode", T: _p0.Func, F: _encode},
}
