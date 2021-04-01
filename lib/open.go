package lib

import (
	"context"
	"os"

	"github.com/dcaiafa/nitro"
)

func open(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	filename, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{&Reader{f}}, nil
}
