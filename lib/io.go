package lib

import (
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type reader struct {
	io.Reader
}

func (r *reader) String() string { return "<Reader>" }
func (r *reader) Type() string   { return "Reader" }

func wrapReader(r io.Reader) nitro.Value {
	v, ok := r.(nitro.Value)
	if ok {
		return v
	}
	return &reader{r}
}

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func in(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{wrapReader(os.Stdin)}, nil
}
