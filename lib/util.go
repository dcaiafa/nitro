package lib

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

var (
	errNotEnoughArgs       = errors.New("not enough arguments")
	errTooManyArgs         = errors.New("too many arguments")
	errInvalidNumberOfArgs = errors.New("invalid number of arguments")
)

func expectArgCount(args []vm.Value, min, max int) error {
  switch {
  case len(args) < min:
    return errNotEnoughArgs
  case len(args) > max:
    return errTooManyArgs
  default:
    return nil
  }
}

func getValueArg(args []vm.Value, ndx int) (nitro.Value, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	return args[ndx], nil
}

func getIntArg(args []vm.Value, ndx int) (int64, error) {
	if ndx >= len(args) {
		return 0, errNotEnoughArgs
	}
	v, ok := args[ndx].(vm.Int)
	if !ok {
		return 0, fmt.Errorf(
			"expected argument %d to be Int, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v.Int64(), nil
}

func getFloatArg(args []vm.Value, ndx int) (float64, error) {
	if ndx >= len(args) {
		return 0, errNotEnoughArgs
	}
	if v, ok := args[ndx].(vm.Float); ok {
    return v.Float64(), nil
  } else if v, ok := args[ndx].(vm.Int); ok {
    return float64(v.Int64()), nil
  } else {
    return 0, errExpectedArg(ndx, args[ndx], "float", "int")
  }
}

func getBoolArg(args []vm.Value, ndx int) (bool, error) {
	if ndx >= len(args) {
		return false, errNotEnoughArgs
	}
	v, ok := args[ndx].(vm.Bool)
	if !ok {
		return false, fmt.Errorf(
			"expected argument %d to be Bool, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v.Bool(), nil
}

func getStringArg(args []vm.Value, ndx int) (string, error) {
	if ndx >= len(args) {
		return "", errNotEnoughArgs
	}
	v, ok := args[ndx].(vm.String)
	if !ok {
		return "", fmt.Errorf(
			"expected argument %d to be String, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v.String(), nil
}

func getListArg(args []vm.Value, ndx int) (*nitro.Array, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*vm.Array)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be list, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getObjectArg(args []vm.Value, ndx int) (*nitro.Object, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*vm.Object)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be Object, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getTimeArg(args []vm.Value, ndx int) (Time, error) {
	if ndx >= len(args) {
		return Time{}, errNotEnoughArgs
	}
	v, ok := args[ndx].(Time)
	if !ok {
		return Time{}, fmt.Errorf(
			"expected argument %d to be time, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getRegexArg(args []vm.Value, ndx int) (*nitro.Regex, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*vm.Regex)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be Regex, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getDurationArg(args []vm.Value, ndx int) (Duration, error) {
	if ndx >= len(args) {
		return Duration{}, errNotEnoughArgs
	}
	v, ok := args[ndx].(Duration)
	if !ok {
		return Duration{}, fmt.Errorf(
			"expected argument %d to be duration, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getProcessArg(args []vm.Value, ndx int) (*process, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*process)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be process, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getFileArg(args []vm.Value, ndx int) (*File, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*File)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be file, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getArg(args []vm.Value, ndx int) (*process, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*process)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be process, but it is %v",
			ndx+1, nitro.TypeName(args[ndx]))
	}
	return v, nil
}

func getWriterArg(args []vm.Value, ndx int) (io.Writer, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	switch v := args[ndx].(type) {
	case io.Writer:
		return v, nil
	default:
		return nil, fmt.Errorf("argument %v is not writable", nitro.TypeName(v))
	}
}

func getIterArg(m *nitro.VM, args []vm.Value, ndx int) (nitro.Iterator, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v := args[ndx]
	it, err := nitro.MakeIterator(m, v)
	if err != nil {
		return nil, fmt.Errorf(
			"argument #%v %v is not iterable", ndx+1, nitro.TypeName(v))
	}
	return it, nil
}

func getReaderArg(vmArg *vm.VM, args []vm.Value, ndx int) (vm.Reader, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v := args[ndx]
	reader, err := vm.MakeReader(vmArg, v)
	if err != nil {
		return nil, fmt.Errorf("argument #%v %v is not readable", ndx+1, nitro.TypeName(v))
	}
	return reader, nil
}

func getCallableArg(args []nitro.Value, ndx int) (nitro.Callable, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v := args[ndx]
	callable, ok := args[ndx].(nitro.Callable)
	if !ok {
		return nil, fmt.Errorf("argument #%v %v is not callable", ndx+1, nitro.TypeName(v))
	}
	return callable, nil
}

func errExpectedArg(ndx int, actual nitro.Value, expected ...string) error {
	return fmt.Errorf(
		"expected argument #%v to be %v, but it was %v",
		ndx+1, strings.Join(expected, " or "),
		nitro.TypeName(actual))
}
