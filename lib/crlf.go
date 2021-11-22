package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

var errFromCRLFUsage = nitro.NewInvalidUsageError("from_crlf(string|reader)")

func fromCRLF(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errFromCRLFUsage
	}

	switch source := args[0].(type) {
	case nitro.String:
		converter := newFromCRLFReader(source.MakeReader())
		converted, err := io.ReadAll(converter)
		if err != nil {
			return nil, err
		}
		return []nitro.Value{nitro.NewString(string(converted))}, nil

	case io.Reader, io.Writer:
		converter := newFromCRLFReader(source)
		return []nitro.Value{converter}, nil

	default:
		return nil, errFromCRLFUsage
	}
}

var errToCRLFUsage = nitro.NewInvalidUsageError("to_crlf(string|reader, bool?)")

func toCRLF(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errToCRLFUsage
	}

	convert := runtime.GOOS == "windows"
	if len(args) == 2 {
		convertArg, ok := args[1].(nitro.Bool)
		if !ok {
			return nil, errToCRLFUsage
		}
		convert = convertArg.Bool()
	}

	switch source := args[0].(type) {
	case nitro.String:
		if !convert {
			return []nitro.Value{source}, nil
		}

		converter := newToCRLFReader(strings.NewReader(source.String()))
		converted, err := io.ReadAll(converter)
		if err != nil {
			return nil, err
		}
		return []nitro.Value{nitro.NewString(string(converted))}, nil

	case io.Reader:
		if !convert {
			return []nitro.Value{args[0]}, nil
		}

		converter := newToCRLFReader(source)
		return []nitro.Value{converter}, nil

	default:
		return nil, errToCRLFUsage
	}
}

type fromCRLFReaderWriter struct {
	nitro.BaseValue

	c   io.Closer
	r   io.Reader
	w   io.Writer
	buf *bufio.Reader
}

func newFromCRLFReader(v nitro.Value) *fromCRLFReaderWriter {
	c := &fromCRLFReaderWriter{
		BaseValue: nitro.BaseValue{TypeName: "fromcrlf"},
	}

	var ok bool
	c.c, _ = v.(io.Closer)
	c.r, _ = v.(io.Reader)
	c.w, _ = v.(io.Writer)

	if c.r != nil {
		c.buf, ok = c.r.(*bufio.Reader)
		if !ok {
			c.buf = bufio.NewReader(c.r)
		}
	}

	return c
}

func (c *fromCRLFReaderWriter) Read(buf []byte) (int, error) {
	if c.r == nil {
		return 0, fmt.Errorf("cannot read from %v", c.TypeName)
	}

	n := 0
	for n < len(buf) {
		b, err := c.buf.ReadByte()
		if err != nil {
			return n, err
		}
		if b != '\r' {
			buf[n] = b
			n++
		}
	}
	return n, nil
}

func (c *fromCRLFReaderWriter) Write(buf []byte) (int, error) {
	if c.w == nil {
		return 0, fmt.Errorf("cannot write to %v", c.TypeName)
	}

	n := 0
	for len(buf) > 0 {
		if buf[0] == '\r' {
			buf = buf[1:]
			n++
		}
		i := bytes.IndexByte(buf, '\r')
		if i == -1 {
			i = len(buf)
		}
		if i > 0 {
			pn, err := c.w.Write(buf[:i])
			n += pn
			if err != nil {
				return n, err
			}
			buf = buf[i:]
		}
	}
	return n, nil
}

func (c *fromCRLFReaderWriter) Call(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, core.ErrWriterCallUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, core.ErrWriterCallUsage
	}

	n, err := io.Copy(c, reader)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewInt(n)}, nil
}

func (c *fromCRLFReaderWriter) Close() error {
	if c.c != nil {
		return c.c.Close()
	}
	return nil
}

type toCRLFReader struct {
	nitro.BaseValue

	r   io.Reader
	buf *bufio.Reader
	cr  bool
}

func newToCRLFReader(r io.Reader) *toCRLFReader {
	c := &toCRLFReader{
		BaseValue: nitro.BaseValue{TypeName: "tocrlf"},
		r:         r,
	}

	var ok bool
	c.buf, ok = r.(*bufio.Reader)
	if !ok {
		c.buf = bufio.NewReader(r)
	}

	return c
}

func (c *toCRLFReader) Read(buf []byte) (int, error) {
	n := 0
	for n < len(buf) {
		b, err := c.readByte()
		if err != nil {
			return n, err
		}
		buf[n] = b
		n++
	}
	return n, nil
}

func (c *toCRLFReader) Close() error {
	if closer, ok := c.r.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (c *toCRLFReader) readByte() (byte, error) {
	b, err := c.buf.ReadByte()
	if err != nil {
		return 0, err
	}
	if b == '\n' && !c.cr {
		c.cr = true
		c.buf.UnreadByte()
		return '\r', nil
	}
	c.cr = (b == '\r')

	return b, nil
}
