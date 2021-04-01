package lib

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro"
)

func fnPrintAll(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	stdout := Stdout(ctx)
	e, err := getEnumeratorArg(ctx, args, 0)
	if err != nil {
		return nil, err
	}
	for {
		v, ok, err := nitro.Next(ctx, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		fmt.Fprintln(stdout, valuesToInterface(v)...)
	}
	return nil, nil
}
