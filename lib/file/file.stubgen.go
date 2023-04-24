package file

import (
	_p0 "github.com/dcaiafa/nitro/internal/export"
	_p1 "github.com/dcaiafa/nitro/internal/stub"

	_p2 "github.com/dcaiafa/nitro/internal/vm"
)

type OpenOptions struct {
	Read   bool
	Write  bool
	Append bool
	Create bool
	Excl   bool
	Trunc  bool
}

func (m *OpenOptions) FromMap(v *_p2.Object) error {
	var err error
	_ = err
	v.ForEach(func(k, v _p2.Value) bool {
		n, ok := k.(_p2.String)
		if !ok {
			err = _p1.ErrMapKeyMustBeStr
			return false
		}
		switch n.String() {
		case "read":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Read = tv
		case "write":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Write = tv
		case "append":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Append = tv
		case "create":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Create = tv
		case "excl":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Excl = tv
		case "trunc":
			cv, ok := v.(_p2.Bool)
			if !ok {
				err = _p1.ErrInvalidFieldType
				return false
			}
			tv := (cv).Bool()
			m.Trunc = tv
		default:
			err = _p1.StructDoesNotHaveField("OpenOptions", n.String())
			return false
		}
		return true
	})
	if err != nil {
		return err
	}
	return nil
}
func _open(vm *_p2.VM, args []_p2.Value, nret int) ([]_p2.Value, error) {
	var err error
	_ = err
	if len(args) < 2 {
		return nil, _p1.ErrInsufficientArgs
	}
	switch _a0 := args[0].(type) {
	case _p2.String:
		if len(args) == 1 {
			{
				_ta0 := (_a0).String()
				_r0, err := open0(vm, _ta0)
				if err != nil {
					return nil, err
				}
				return []_p2.Value{_r0}, nil
			}
		}
		switch _a1 := args[1].(type) {
		case *_p2.Object:
			if len(args) > 2 {
				return nil, _p1.ErrTooManyArgs
			}
			{
				_ta0 := (_a0).String()
				var _ta1 *OpenOptions
				if _a1 != nil {
					_ta1 = new(OpenOptions)
					err = _ta1.FromMap(_a1)
					if err != nil {
						return nil, _p1.InvalidArgErr(args, 2, err)
					}
				}
				_r0, err := open1(vm, _ta0, _ta1)
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

var Exports = _p0.Exports{
	{N: "open", T: _p0.Func, F: _open},
}
