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
			"expected argument %d to be String, but it is %v",
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

func getWriterArg(args []runtime.Value, ndx int) (io.Writer, error) {
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

type enumReader struct {
	ctx context.Context
	e   nitro.Value
	buf ByteQueue
}

func (r *enumReader) String() string { return "EnumReader" }
func (r *enumReader) Type() string   { return "EnumReader" }

func (r *enumReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	for len(r.buf.Peek()) < len(b) {
		v, ok, err := nitro.Next(r.ctx, r.e, 1)
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		str, ok := v[0].(nitro.String)
		if !ok {
			return 0, fmt.Errorf(
				"cannot stream enumerator of %q",
				nitro.TypeName(v[0]))
		}
		r.buf.Write([]byte(str.String()))
		r.buf.Write([]byte{'\n'})
	}

	if len(r.buf.Peek()) == 0 {
		return 0, io.EOF
	}

	n := len(r.buf.Peek())
	if n > len(b) {
		n = len(b)
	}
	copy(b, r.buf.Peek()[:n])
	r.buf.Pop(n)
	return n, nil
}

func ToReader(ctx context.Context, v runtime.Value) (io.Reader, error) {
	switch v := v.(type) {
	case io.Reader:
		return v, nil

	case nitro.String:
		return strings.NewReader(v.String()), nil

	case *nitro.Closure:
		return &enumReader{
			ctx: ctx,
			e:   v,
		}, nil

	default:
		return nil, fmt.Errorf(
			"value of type %q is not streamable",
			nitro.TypeName(v))
	}
}
