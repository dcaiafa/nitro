package lib

import (
	"bytes"
	"errors"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

type Buffer struct {
	core.WriterBase
	buf *bytes.Buffer
}

func (b *Buffer) Read(buf []byte) (int, error) {
	return b.buf.Read(buf)
}

func (b *Buffer) String() string {
	return b.buf.String()
}

func NewBuffer(data string) *Buffer {
	b := &Buffer{}
	if data != "" {
		b.buf = bytes.NewBufferString(data)
	} else {
		b.buf = new(bytes.Buffer)
	}
	b.WriterBase = core.NewWriterBase("buffer", b.buf)
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
