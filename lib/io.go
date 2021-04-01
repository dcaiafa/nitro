package lib

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dcaiafa/nitro"
)

type Reader struct {
	io.Reader
}

func (r *Reader) String() string { return "<Reader>" }
func (r *Reader) Type() string   { return "Reader" }

type Writer struct {
	io.Writer
}

func (w *Writer) String() string { return "<Reader>" }
func (w *Writer) Type() string   { return "Reader" }

var stdoutContextKey = "stdout"

func ContextWithStdout(ctx context.Context, stdout io.Writer) context.Context {
	return context.WithValue(ctx, &stdoutContextKey, stdout)
}

func Stdout(ctx context.Context) io.Writer {
	stdout, ok := ctx.Value(&stdoutContextKey).(io.Writer)
	if !ok {
		return os.Stdout
	}
	return stdout
}

func ToReader(ctx context.Context, v nitro.Value) (io.Reader, error) {
	switch v := v.(type) {
	case io.Reader:
		return v, nil
	case nitro.String:
		return strings.NewReader(v.String()), nil
	default:
		return nil, fmt.Errorf("Value %v is not readable", nitro.TypeName(v))
	}
}

func in(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Reader{os.Stdin}}, nil
}

func out(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Writer{Stdout(ctx)}}, nil
}

func print(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	iargs := valuesToInterface(args)
	fmt.Fprintln(Stdout(ctx), iargs...)
	return nil, nil
}

func valuesToInterface(values []nitro.Value) []interface{} {
	ivalues := make([]interface{}, len(values))
	for i, v := range values {
		switch v := v.(type) {
		case nitro.Int:
			ivalues[i] = v.Int64()
		case nitro.Float:
			ivalues[i] = v.Float64()
		case nitro.String:
			ivalues[i] = v.String()
		default:
			ivalues[i] = v
		}
	}
	return ivalues
}
