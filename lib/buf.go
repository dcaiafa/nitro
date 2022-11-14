package lib

import (
	"bytes"
	"io"

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

func (b *Buffer) Len() int {
	return b.buf.Len()
}

func (b *Buffer) String() string {
	return b.buf.String()
}

func newBuffer(data string) *Buffer {
	b := &Buffer{}
	if data != "" {
		b.buf = bytes.NewBufferString(data)
	} else {
		b.buf = new(bytes.Buffer)
	}
	b.WriterBase = core.NewWriterBase("buffer", b.buf)
	return b
}

func getBufferArg(args []nitro.Value, ndx int) (*Buffer, error) {
	if ndx >= len(args) {
		return nil, errNotEnoughArgs
	}
	v, ok := args[ndx].(*Buffer)
	if !ok {
		return nil, errExpectedArg(ndx, args[ndx], "duration")
	}
	return v, nil
}

func bufNew(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	if err = expectArgCount(args, 0, 1); err != nil {
		return nil, err
	}

	init := ""
	if len(args) == 1 {
		init, err = getStringArg(args, 0)
		if err != nil {
			return nil, err
		}
	}

	b := newBuffer(init)

	return []nitro.Value{b}, nil
}

func bufRead(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return read(m, args, nRet)
}

func bufReadByte(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	b, err := buf.buf.ReadByte()
	if err != nil {
		if err == io.EOF {
			return []nitro.Value{nil}, nil
		}
	}

	return []nitro.Value{nitro.NewInt(int64(b))}, nil
}

func bufReadRune(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	r, l, err := buf.buf.ReadRune()
	if err != nil {
		if err == io.EOF {
			return []nitro.Value{nil}, nil
		}
	}

	return []nitro.Value{nitro.NewInt(int64(r)), nitro.NewInt(int64(l))}, nil
}

func bufReadFrom(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	r, err := getReaderArg(m, args, 1)
	if err != nil {
		return nil, err
	}

	_, err = buf.buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func bufLen(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	l := buf.Len()

	return []nitro.Value{nitro.NewInt(int64(l))}, nil
}

func bufCap(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	c := buf.buf.Cap()

	return []nitro.Value{nitro.NewInt(int64(c))}, nil
}

func bufUnreadByte(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	buf, err := getBufferArg(args, 0)
	if err != nil {
		return nil, err
	}

	err = buf.buf.UnreadByte()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
