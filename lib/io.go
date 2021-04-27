package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/dcaiafa/nitro"
)

type reader struct {
	io.Reader
}

func (r *reader) String() string { return "<reader>" }
func (r *reader) Type() string   { return "reader" }

func (r *reader) EvalBinOp(op nitro.BinOp, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("reader does not support this operation")
}

func (r *reader) EvalUnaryMinus() (nitro.Value, error) {
	return nil, fmt.Errorf("reader does not support this operation")
}

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

func in(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return []nitro.Value{wrapReader(os.Stdin)}, nil
}
