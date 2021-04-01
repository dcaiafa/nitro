package lib

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro"
)

func fnPrint(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, expRetN int) ([]nitro.Value, error) {
	stdout := Stdout(ctx)
	iargs := valuesToInterface(args)
	fmt.Fprintln(stdout, iargs...)
	return nil, nil
}

func valuesToInterface(values []nitro.Value) []interface{} {
	ivalues := make([]interface{}, len(values))
	for i, v := range values {
		switch v := v.(type) {
		case nitro.Int:
			ivalues[i] = v.Int64()
		case nitro.Float:
			ivalues[i] = v.Float64()
		case nitro.String:
			ivalues[i] = v.String()
		default:
			ivalues[i] = v
		}
	}
	return ivalues
}
