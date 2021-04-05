package lib

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
)

func lines(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}
	input, err := ToReader(ctx, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument #1: %w", err)
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
		CloseReader(l.input)
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
