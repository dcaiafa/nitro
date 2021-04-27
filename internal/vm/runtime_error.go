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

func (e *RuntimeError) EvalBinOp(op BinOp, operand Value) (Value, error) {
	return nil, fmt.Errorf("error does not support this operation")
}

func (e *RuntimeError) EvalUnaryMinus() (Value, error) {
	return nil, fmt.Errorf("error does not support this operation")
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

func (e *RuntimeError) String() string {
	return e.Error()
}

func (e *RuntimeError) Type() string {
	return "RuntimeError"
}

func (e *RuntimeError) Error() string {
	str := strings.Builder{}
	if e.Err != nil {
		str.WriteString(e.Err.Error())
	} else if e.ErrValue != nil {
		str.WriteString(e.ErrValue.String())
	}

	for _, f := range e.Stack {
		var loc string
		if f.Filename != "" {
			loc = fmt.Sprintf("%v:%v", f.Filename, f.Line)
		} else {
			loc = "<builtin>"
		}
		fmt.Fprintf(&str, "\n %v  %v", loc, f.Func)
	}

	return str.String()
}

func (e *RuntimeError) Unwrap() error {
	return e.Err
}

var ErrCannotCallNil = errors.New("cannot evaluate function call because target is nil")
var ErrIsNotIterable = errors.New("is not iterable")
