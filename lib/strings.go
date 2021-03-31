package lib

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/dcaiafa/nitro"
)

func lines(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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

	l := &linesState{
		input:   input,
		scanner: bufio.NewScanner(input),
	}

	closure := nitro.NewClosure(l.Next, nil)

	return []nitro.Value{closure}, nil
}

type linesState struct {
	input   io.Reader
	scanner *bufio.Scanner
}

func (l *linesState) Next(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if !l.scanner.Scan() {
		if l.scanner.Err() != nil {
			return nil, l.scanner.Err()
		}
		return []nitro.Value{
			nitro.NewBool(false),
			nil,
		}, nil
	}
	return []nitro.Value{
		nitro.NewBool(true),
		nitro.NewString(l.scanner.Text()),
	}, nil
}
