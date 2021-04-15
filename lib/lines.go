package lib

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
)

func lines(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}
	input, err := ToReader(m, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument #1: %w", err)
	}

	l := &linesIter{
		input:   input,
		scanner: bufio.NewScanner(input),
	}

	outIter := nitro.NewIterator(l.Next, nil, 1)
	return []nitro.Value{outIter}, nil
}

type linesIter struct {
	input   io.Reader
	scanner *bufio.Scanner
}

func (l *linesIter) Next(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
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
