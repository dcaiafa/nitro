package lib

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type stdinWrapper struct {
	nitro.BaseValue
	io.Reader
}

func wrapStdin(r io.Reader) nitro.Value {
	v, ok := r.(nitro.Value)
	if ok {
		return v
	}
	return &stdinWrapper{
		BaseValue: nitro.BaseValue{TypeName: "reader"},
		Reader:    r,
	}
}

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func stdin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{wrapStdin(os.Stdin)}, nil
}

type writerBase struct {
	io.Writer
	typ string
}

func newWriterBase(typ string, w io.Writer) writerBase {
	return writerBase{Writer: w, typ: typ}
}

func (b *writerBase) String() string { return "<" + b.Type() + ">" }
func (b *writerBase) Type() string   { return b.typ }

func (b *writerBase) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("<writer> does not support this operation")
}

var errWriterCallUsage = errors.New(
	`invalid usage. Expected <writer>(reader)`)

func (b *writerBase) Call(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errWriterCallUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, errWriterCallUsage
	}

	n, err := io.Copy(b.Writer, reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}

func wrapWriter(typ string, w io.Writer) nitro.Value {
	v, ok := w.(nitro.Value)
	if ok {
		return v
	}
	wb := newWriterBase(typ, w)
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
			return outStack.stack[len(outStack.stack)-1].(io.Writer)
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
		return nil, errWriterCallUsage
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
		return nil, errWriterCallUsage
	}

	n, err := io.Copy(Stderr(m), reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}
