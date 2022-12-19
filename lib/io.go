package lib

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
)

type stdinWrapper struct {
	*os.File
}

func (w *stdinWrapper) String() string             { return "<stdin>" }
func (w *stdinWrapper) Type() string               { return "stdin" }
func (w *stdinWrapper) Traits() vm.Traits          { return vm.TraitNone }
func (w *stdinWrapper) GetNativeReader() io.Reader { return w.File }

var Stdin = &stdinWrapper{
	File: os.Stdin,
}

func stdin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{Stdin}, nil
}

func wrapWriter(typ string, w io.Writer) nitro.Value {
	v, ok := w.(nitro.Value)
	if ok {
		return v
	}
	wb := core.NewWriterBase(typ, w)
	return &wb
}

var stdoutUserDataKey = "stdout"

type stdoutStack struct {
	stack []io.Writer
}

func Stdout(m *nitro.VM) io.Writer {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if ok {
		if len(outStack.stack) != 0 {
			return outStack.stack[len(outStack.stack)-1]
		}
	}
	return os.Stdout
}

func SetStdout(m *nitro.VM, w io.Writer) {
	outStack := &stdoutStack{}
	outStack.stack = []io.Writer{w}
	m.SetUserData(&stdoutUserDataKey, outStack)
}

func PushStdout(m *nitro.VM, out io.Writer) {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if !ok {
		outStack = &stdoutStack{}
		m.SetUserData(&stdoutUserDataKey, outStack)
	}
	outStack.stack = append(outStack.stack, out)
}

func PopStdout(m *nitro.VM) io.Writer {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if !ok || len(outStack.stack) == 0 {
		return nil
	}
	prevOut := outStack.stack[len(outStack.stack)-1]
	outStack.stack[len(outStack.stack)-1] = nil
	outStack.stack = outStack.stack[:len(outStack.stack)-1]
	return prevOut
}

var errOutUsage = errors.New(
	`invalid usage. Expected out(reader?)`)

func stdout(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return []nitro.Value{wrapWriter("stdout", Stdout(m))}, nil
	}

	if len(args) != 1 {
		return nil, errOutUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, core.ErrWriterCallUsage
	}

	n, err := io.Copy(Stdout(m), reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}

func pushStdout(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	out, err := getWriterArg(args, 0)
	if err != nil {
		return nil, err
	}
	PushStdout(m, out)
	return nil, nil
}

func popStdout(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	prevOut := PopStdout(m)
	if prevOut == nil {
		return nil, fmt.Errorf("the stdout stack is empty")
	}
	return []nitro.Value{wrapWriter("stdout", prevOut)}, nil
}

var stderrUserDataKey = "stderr"

func Stderr(m *nitro.VM) io.Writer {
	w, ok := m.GetUserData(&stdoutUserDataKey).(io.Writer)
	if ok {
		return w
	}
	return os.Stderr
}

func SetStderr(m *nitro.VM, w io.Writer) {
	m.SetUserData(&stderrUserDataKey, w)
}

var errErrUsage = errors.New(
	`invalid usage. Expected err(reader?)`)

func stderr(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return []nitro.Value{wrapWriter("err", Stderr(m))}, nil
	}

	if len(args) != 1 {
		return nil, errErrUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, core.ErrWriterCallUsage
	}

	n, err := io.Copy(Stderr(m), reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}
