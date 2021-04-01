package lib

import (
	"context"
	"io"
	"os"

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

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func in(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Reader{os.Stdin}}, nil
}

func out(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Writer{Stdout(ctx)}}, nil
}
