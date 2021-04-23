package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type writer struct {
	io.Writer
}

func (w *writer) String() string { return "<writer>" }
func (w *writer) Type() string   { return "writer" }

func (w *writer) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("writer does not support this operation")
}

func (w *writer) EvalUnaryMinus() (nitro.Value, error) {
	return nil, fmt.Errorf("writer does not support this operation")
}

func wrapWriter(w io.Writer) nitro.Value {
	v, ok := w.(nitro.Value)
	if ok {
		return v
	}
	return &writer{w}
}

var stdoutUserDataKey = "stdout"

type stdoutStack struct {
	stack []io.Writer
}

func Stdout(m *nitro.Machine) io.Writer {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
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

func PushOut(m *nitro.Machine, out io.Writer) {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if !ok {
		outStack = &stdoutStack{}
		m.SetUserData(&stdoutUserDataKey, outStack)
	}
	outStack.stack = append(outStack.stack, out)
}

func PopOut(m *nitro.Machine) io.Writer {
	outStack, ok := m.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if !ok || len(outStack.stack) == 0 {
		return nil
	}
	prevOut := outStack.stack[len(outStack.stack)-1]
	outStack.stack[len(outStack.stack)-1] = nil
	outStack.stack = outStack.stack[:len(outStack.stack)-1]
	return prevOut
}

func out(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{wrapWriter(Stdout(m))}, nil
}

func pushout(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	out, err := getWriterArg(args, 0)
	if err != nil {
		return nil, err
	}
	PushOut(m, out)
	return nil, nil
}

func popout(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	prevOut := PopOut(m)
	if prevOut == nil {
		return nil, fmt.Errorf("the stdout stack is empty")
	}
	return []nitro.Value{wrapWriter(prevOut)}, nil
}
