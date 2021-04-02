package lib

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/runtime"
)

func getIntArg(args []runtime.Value, ndx int) (int64, error) {
	if ndx >= len(args) {
		return 0, errNotEnoughArgs
	}
	v, ok := args[ndx].(runtime.Int)
	if !ok {
		return 0, fmt.Errorf(
			"expected argument %d to be Int, but it is %v",
			ndx+1, args[ndx].Type())
	}
	return v.Int64(), nil
}

func getStringArg(args []runtime.Value, ndx int) (string, error) {
	if ndx >= len(args) {
		return "", errNotEnoughArgs
	}
	v, ok := args[ndx].(runtime.String)
	if !ok {
		return "", fmt.Errorf(
			"expected argument %d to be Int, but it is %v",
			ndx+1, args[ndx].Type())
	}
	return v.String(), nil
}

func getObjectArg(args []runtime.Value, ndx int) (*nitro.Object, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*runtime.Object)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be Object, but it is %v",
			ndx+1, args[ndx].Type())
	}
	return v, nil
}

func getRegexArg(args []runtime.Value, ndx int) (*nitro.Regex, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*runtime.Regex)
	if !ok {
		return nil, fmt.Errorf(
			"expected argument %d to be Regex, but it is %v",
			ndx+1, args[ndx].Type())
	}
	return v, nil
}

func getReaderArg(args []runtime.Value, ndx int) (io.Reader, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}

	switch v := args[ndx].(type) {
	case io.Reader:
		return v, nil
	case nitro.String:
		return strings.NewReader(v.String()), nil
	default:
		return nil, fmt.Errorf("argument %v is not readable", nitro.TypeName(v))
	}
}

func getEnumeratorArg(ctx context.Context, args []runtime.Value, ndx int) (nitro.Value, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}

	v, err := nitro.MakeEnumerator(ctx, args[ndx])
	if err != nil {
		return nil, fmt.Errorf("argument %v is not enumerable: %v", args[ndx], err)
	}

	return v, nil
}

func getCallableArg(ctx context.Context, args []runtime.Value, ndx int) (nitro.Value, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}

	switch v := args[ndx].(type) {
	case *nitro.Closure, nitro.ExternFn, *nitro.Func:
		return v, nil

	default:
		return nil, fmt.Errorf("argument %v is not callable", args[ndx])
	}
}
