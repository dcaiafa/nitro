package vm

import (
	"errors"
	"fmt"
	"strings"
)

type RuntimeError struct {
	Err      error
	ErrValue Value
	Stack    []FrameInfo
}

var _ error = (*RuntimeError)(nil)

func wrapRuntimeError(m *VM, err *error) *RuntimeError {
	if *err == nil {
		return nil
	}
	var rerr *RuntimeError
	if !errors.As(*err, &rerr) {
		rerr = &RuntimeError{
			Err: *err,
		}
	}
	if rerr.Stack == nil {
		rerr.Stack = m.GetStackInfo()
	}
	*err = rerr
	return rerr
}

func (e *RuntimeError) Index(k Value) (Value, error) {
	k, ok := k.(String)
	if !ok {
		return nil, fmt.Errorf("RuntimeError does not have field %v", k)
	}

	switch k.String() {
	case "error":
		errValue := e.ErrValue
		if errValue == nil {
			errValue = NewString(e.Err.Error())
		}
		return errValue, nil

	default:
		return nil, fmt.Errorf("RuntimeError does not have field %v", k)
	}
}

func (e *RuntimeError) IndexRef(k Value) (ValueRef, error) {
	return ValueRef{}, fmt.Errorf("RuntimeError is read-only")
}

func (e *RuntimeError) String() string { return e.Error() }
func (e *RuntimeError) Type() string   { return "error" }
func (e *RuntimeError) Traits() Traits { return TraitNone }

func (e *RuntimeError) Message() string {
	if e.Err != nil {
		return e.Err.Error()
	} else if e.ErrValue != nil {
		return e.ErrValue.String()
	} else {
		return "<no message>"
	}
}

func (e *RuntimeError) Error() string {
	str := strings.Builder{}
	str.WriteString(e.Message())

	for _, f := range e.Stack {
		fmt.Fprintf(&str, "\n %v", f.String())
	}

	return str.String()
}

func (e *RuntimeError) Unwrap() error {
	return e.Err
}

var ErrCannotCallNil = errors.New("cannot evaluate function call because target is nil")
var ErrIsNotIterable = errors.New("is not iterable")
var ErrIsNotReadable = errors.New("is not readable")
var ErrDivideByZero = errors.New("divide by zero")

type InvalidUsageError struct {
	expected string
}

func NewInvalidUsageError(expected string) *InvalidUsageError {
	return &InvalidUsageError{expected: expected}
}

func (e *InvalidUsageError) Error() string {
	return "invalid usage. Expected " + e.expected
}
