package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type stdin struct {
	io.Reader
}

func (r *stdin) String() string { return "<reader>" }
func (r *stdin) Type() string   { return "reader" }

func (r *stdin) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("reader does not support this operation")
}

func wrapStdin(r io.Reader) nitro.Value {
	v, ok := r.(nitro.Value)
	if ok {
		return v
	}
	return &stdin{r}
}

func CloseReader(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func in(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{wrapStdin(os.Stdin)}, nil
}
