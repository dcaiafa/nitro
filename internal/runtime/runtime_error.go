package runtime

import (
	"errors"
	"fmt"
	"strings"
)

type RuntimeError struct {
	Err   error
	Stack []FrameInfo
}

var _ error = (*RuntimeError)(nil)

func (e *RuntimeError) Error() string {
	str := strings.Builder{}
	str.WriteString("Runtime error: " + e.Err.Error())
	for _, f := range e.Stack {
		fmt.Fprintf(&str, "\n %v:%v", f.Filename, f.Line)
	}
	return str.String()
}

func (e *RuntimeError) Unwrap() error {
	return e.Err
}

var ErrCannotCallNil = errors.New("cannot evaluate function call because target is nil")
