package core

import (
	"errors"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

type NativeReader interface {
	GetNativeReader() io.Reader
}

type NativeWriter interface {
	GetNativeWriter() io.Writer
}

type WriterBase struct {
	io.Writer
	typ string
}

func NewWriterBase(w io.Writer) WriterBase {
	return WriterBase{Writer: w}
}

func (b *WriterBase) String() string    { return "<Writer>" }
func (b *WriterBase) Type() string      { return "Writer" }
func (b *WriterBase) Traits() vm.Traits { return vm.TraitNone }

var ErrWriterCallUsage = errors.New(
	`invalid usage. Expected <writer>(reader)`)

func (b *WriterBase) Call(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, ErrWriterCallUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, ErrWriterCallUsage
	}

	n, err := io.Copy(b.Writer, reader)
	if err != nil {
		return nil, err
	}

	CloseReader(reader)

	return []nitro.Value{nitro.NewInt(n)}, nil
}

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

var ErrNotEnoughArgs = errors.New("not enough arguments")
var ErrTooManyArgs = errors.New("too many arguments")

func CheckArgCount(args []nitro.Value, min, max int) error {
	if len(args) < min {
		return ErrNotEnoughArgs
	} else if len(args) > max {
		return ErrTooManyArgs
	}
	return nil
}

func GetString(args []nitro.Value, n int) (nitro.String, error) {
	arg, ok := args[n].(nitro.String)
	if !ok {
		return nitro.String{}, fmt.Errorf(
			"arg %v: expected string, got %v",
			n, nitro.TypeName(args[n]))
	}
	return arg, nil
}

func GetInt(args []nitro.Value, n int) (nitro.Int, error) {
	arg, ok := args[n].(nitro.Int)
	if !ok {
		return nitro.Int{}, fmt.Errorf(
			"arg %v: expected int, got %v",
			n, nitro.TypeName(args[n]))
	}
	return arg, nil
}

func GetFloat(args []nitro.Value, n int) (nitro.Float, error) {
	arg, ok := args[n].(nitro.Float)
	if !ok {
		return nitro.Float{}, fmt.Errorf(
			"arg %v: expected float, got %v",
			n, nitro.TypeName(args[n]))
	}
	return arg, nil
}

func GetReader(vm *nitro.VM, args []nitro.Value, n int) (nitro.Reader, error) {
	arg, err := nitro.MakeReader(vm, args[n])
	if err != nil {
		return nil, fmt.Errorf("arg %v: %w", n, err)
	}
	return arg, nil
}
