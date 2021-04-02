package lib

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type Writer struct {
	io.Writer
}

func (w *Writer) String() string { return "<Writer>" }
func (w *Writer) Type() string   { return "Writer" }

var stdoutUserDataKey = "stdout"

type stdoutStack struct {
	stack []io.Writer
}

func Stdout(ctx context.Context) io.Writer {
	outStack, ok := nitro.GetUserData(ctx, &stdoutUserDataKey).(*stdoutStack)
	if ok {
		if len(outStack.stack) != 0 {
			return outStack.stack[len(outStack.stack)-1].(io.Writer)
		}
	}
	return os.Stdout
}

func SetStdout(m *nitro.Machine, w io.Writer) {
	outStack := &stdoutStack{}
	outStack.stack = []io.Writer{w}
	m.SetUserData(&stdoutUserDataKey, outStack)
}

func PushOut(ctx context.Context, out io.Writer) {
	outStack, ok := nitro.GetUserData(ctx, &stdoutUserDataKey).(*stdoutStack)
	if !ok {
		outStack = &stdoutStack{}
		nitro.SetUserData(ctx, &stdoutUserDataKey, outStack)
	}
	outStack.stack = append(outStack.stack, out)
}

func PopOut(ctx context.Context) error {
	outStack, ok := nitro.GetUserData(ctx, &stdoutUserDataKey).(*stdoutStack)
	if !ok || len(outStack.stack) == 0 {
		return fmt.Errorf("the stdout stack is empty")
	}
	outStack.stack[len(outStack.stack)-1] = nil
	outStack.stack = outStack.stack[:len(outStack.stack)-1]
	return nil
}

func fnOut(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Writer{Stdout(ctx)}}, nil
}

func fnPushOut(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	out, err := getWriterArg(args, 0)
	if err != nil {
		return nil, err
	}
	PushOut(ctx, out)
	return nil, nil
}

func fnPopOut(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	err := PopOut(ctx)
	return nil, err
}
