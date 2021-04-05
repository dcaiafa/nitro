package lib

import (
	"context"
	"strings"

	"github.com/dcaiafa/nitro"
)

func fnTrim(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	s, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	res := strings.TrimSpace(s)

	return []nitro.Value{nitro.NewString(res)}, nil
}
