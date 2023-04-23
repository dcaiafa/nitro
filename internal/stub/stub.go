package stub

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro/internal/vm"
)

var (
	ErrInsufficientArgs = errors.New("not enough arguments")
	ErrTooManyArgs      = errors.New("too many arguments")
	ErrInvalidArg       = errors.New("invalid argument")
	ErrMapKeyMustBeStr  = errors.New("map key must be Str")
	ErrInvalidFieldType = errors.New("invalid field type")
)

func InvalidArg(args []vm.Value, idx int) error {
	return fmt.Errorf("%w at index %v", ErrInvalidArg, idx)
}

func InvalidArgErr(args []vm.Value, idx int, err error) error {
	return fmt.Errorf("%w at index %v: %v", ErrInvalidArg, idx, err)
}

func StructDoesNotHaveField(structName, fieldName string) error {
	return fmt.Errorf("%q doesn't have field %q", structName, fieldName)
}

func MustMakeReader(m *vm.VM, readable vm.Value) vm.Reader {
	reader, err := vm.MakeReader(m, readable)
	if err != nil {
		panic(err)
	}
	return reader
}
