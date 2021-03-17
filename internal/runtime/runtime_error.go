package runtime

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

func (e *RuntimeError) String() string {
	return e.Error()
}

func (e *RuntimeError) Type() string {
	return "RuntimeError"
}

func (e *RuntimeError) Error() string {
	str := strings.Builder{}
	str.WriteString("Runtime error: ")
	if e.Err != nil {
		str.WriteString(e.Err.Error())
	} else {
		str.WriteString(e.ErrValue.String())
	}

	for _, f := range e.Stack {
		fmt.Fprintf(&str, "\n %v:%v", f.Filename, f.Line)
	}
	return str.String()
}

func (e *RuntimeError) Unwrap() error {
	return e.Err
}

var ErrCannotCallNil = errors.New("cannot evaluate function call because target is nil")
