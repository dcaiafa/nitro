package io

import (
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
)

//go:generate go run ../../internal/stub/stubgen io.stubgen

type stdinWrapper struct {
	*os.File
}

func (w *stdinWrapper) String() string             { return "<Reader>" }
func (w *stdinWrapper) Type() string               { return "Reader" }
func (w *stdinWrapper) Traits() vm.Traits          { return vm.TraitNone }
func (w *stdinWrapper) GetNativeReader() io.Reader { return w.File }

var DefaultStdin = &stdinWrapper{
	File: os.Stdin,
}

type stdoutWrapper struct {
	*os.File
}

func (w *stdoutWrapper) String() string             { return "<Writer>" }
func (w *stdoutWrapper) Type() string               { return "Writer" }
func (w *stdoutWrapper) Traits() vm.Traits          { return vm.TraitNone }
func (w *stdoutWrapper) GetNativeWriter() io.Writer { return w.File }

var DefaultOut = &stdoutWrapper{
	File: os.Stdout,
}

var DefaultErr = &stdoutWrapper{
	File: os.Stderr,
}

func wrapWriter(w io.Writer) vm.Writer {
	v, ok := w.(vm.Writer)
	if ok {
		return v
	}
	wb := core.NewWriterBase(w)
	return &wb
}

var stdoutUserDataKey = "stdout"

type stdoutStack struct {
	def   vm.Writer
	stack []vm.Writer
}

func getOutState(vm *nitro.VM) *stdoutStack {
	state, ok := vm.GetUserData(&stdoutUserDataKey).(*stdoutStack)
	if !ok {
		state = &stdoutStack{
			def: DefaultOut,
		}
		vm.SetUserData(&stdoutUserDataKey, state)
	}
	return state
}

func Stdout(vm *nitro.VM) vm.Writer {
	state := getOutState(vm)
	if len(state.stack) == 0 {
		return state.def
	}
	return state.stack[len(state.stack)-1]
}

func SetStdout(vm *vm.VM, w io.Writer) {
	state := getOutState(vm)
	state.def = wrapWriter(w)
}

func Stderr(m *nitro.VM) vm.Writer {
	return DefaultErr
}

func in0(vm *vm.VM) (vm.Reader, error) {
	return DefaultStdin, nil
}

func out0(vm *nitro.VM, r vm.Reader) (vm.Writer, error) {
	out := Stdout(vm)
	if r == nil {
		return out, nil
	}
	_, err := io.Copy(out, r)
	if err != nil {
		return nil, err
	}
	core.CloseReader(r)
	return out, nil
}

func err0(vm *nitro.VM, r vm.Reader) (vm.Writer, error) {
	out := Stderr(vm)
	if r == nil {
		return out, nil
	}
	_, err := io.Copy(out, r)
	if err != nil {
		return nil, err
	}
	core.CloseReader(r)
	return out, nil
}

func push_out0(vm *nitro.VM, w vm.Writer) error {
	state := getOutState(vm)
	state.stack = append(state.stack, w)
	return nil
}

func pop_out0(vm *nitro.VM) (vm.Writer, error) {
	state := getOutState(vm)
	if len(state.stack) == 0 {
		return nil, fmt.Errorf("output stack is empty")
	}
	popped := state.stack[len(state.stack)-1]
	state.stack = state.stack[:len(state.stack)-1]
	return popped, nil
}
