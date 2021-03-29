package lib

import (
	"context"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type Reader struct {
	io.Reader
}

func (r *Reader) String() string { return "<Reader>" }
func (r *Reader) Type() string   { return "Reader" }

func In(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	return []nitro.Value{&Reader{os.Stdin}}, nil
}
