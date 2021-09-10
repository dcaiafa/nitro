package lib

import (
	"bufio"
	"io"
	"runtime"
	"strings"

	"github.com/dcaiafa/nitro"
)

var errFromCRLFUsage = nitro.NewInvalidUsageError("fromcrlf(string|reader)")

func fromcrlf(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errFromCRLFUsage
	}

	switch source := args[0].(type) {
	case nitro.String:
		converter := newFromCRLFReader(strings.NewReader(source.String()))
		converted, err := io.ReadAll(converter)
		if err != nil {
			return nil, err
		}
		return []nitro.Value{nitro.NewString(string(converted))}, nil

	case io.Reader:
		converter := newFromCRLFReader(source)
		return []nitro.Value{converter}, nil

	default:
		return nil, errFromCRLFUsage
	}
}

var errToCRLFUsage = nitro.NewInvalidUsageError("tocrlf(string|reader, bool?)")

func tocrlf(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
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

type fromCRLFReader struct {
	nitro.BaseValue

	r   io.Reader
	buf *bufio.Reader
}

func newFromCRLFReader(r io.Reader) *fromCRLFReader {
	c := &fromCRLFReader{
		BaseValue: nitro.BaseValue{TypeName: "fromcrlf"},
		r:         r,
	}

	var ok bool
	c.buf, ok = r.(*bufio.Reader)
	if !ok {
		c.buf = bufio.NewReader(r)
	}

	return c
}

func (c *fromCRLFReader) Read(buf []byte) (int, error) {
	n := 0
	for n < len(buf) {
		b, err := c.readByte(false)
		if err != nil {
			return n, err
		}
		buf[n] = b
		n++
	}
	return n, nil
}

func (c *fromCRLFReader) readByte(cr bool) (byte, error) {
	b, err := c.buf.ReadByte()
	if err == io.EOF && cr {
		return '\r', nil
	} else if err != nil {
		return 0, err
	}
	if b == '\r' && !cr {
		return c.readByte(true)
	}
	if b != '\n' && cr {
		c.buf.UnreadByte()
		return '\r', nil
	}
	return b, nil
}

func (c *fromCRLFReader) Close() error {
	if closer, ok := c.r.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

type fromCRLFWriter struct {
	nitro.BaseValue
	w io.Writer
}

func (c *fromCRLFWriter) Write(buf []byte) (int, error) {
	panic("not implemented")
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
