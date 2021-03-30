package lib

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/dcaiafa/nitro"
)

func Lines(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	var input io.Reader
	switch arg := args[0].(type) {
	case io.Reader:
		input = arg
	case nitro.String:
		input = strings.NewReader(arg.String())
	default:
		return nil, fmt.Errorf("don't know how to split %q into lines", args[0].Type())
	}

	l := &lines{
		input:   input,
		scanner: bufio.NewScanner(input),
	}

	closure := nitro.NewClosure(l.Next, nil)

	return []nitro.Value{closure}, nil
}

type lines struct {
	input   io.Reader
	scanner *bufio.Scanner
}

func (l *lines) Next(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return nil, nil
}
