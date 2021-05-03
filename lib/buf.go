package lib

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) String() string { return b.Buffer.String() }
func (b *Buffer) Type() string   { return "buffer" }

func (b *Buffer) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, fmt.Errorf("buffer does not support this operation")
}

func NewBuffer(data string) *Buffer {
	b := &Buffer{}
	if data != "" {
		b.Buffer = bytes.NewBufferString(data)
	} else {
		b.Buffer = new(bytes.Buffer)
	}
	return b
}

var errBufUsage = errors.New(`invalid usage. Expected buf(string?)`)

func buf(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) > 1 {
		return nil, errBufUsage
	}

	init := ""
	if len(args) == 1 {
		initArg, ok := args[0].(nitro.String)
		if !ok {
			return nil, errBufUsage
		}
		init = initArg.String()
	}

	b := NewBuffer(init)

	return []nitro.Value{b}, nil
}
