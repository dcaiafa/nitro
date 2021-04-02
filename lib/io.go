package lib

import (
	"context"
	"os"

	"github.com/dcaiafa/nitro"
)

func in(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Reader{os.Stdin}}, nil
}
