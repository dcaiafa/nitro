package vm

import (
	"fmt"
	"io"

	"github.com/dcaiafa/nitro/internal/bytequeue"
)

type Readable interface {
	Value
	MakeReader() Reader
}

func IsReadable(v Value) bool {
	if v == nil {
		return true
	}
	switch v.(type) {
	case Readable, Reader, Iterator:
		return true
	default:
		return false
	}
}

func MakeReader(vm *VM, v Value) (Reader, error) {
	if v == nil {
		return emptyReader, nil
	}
	switch v := v.(type) {
	case Reader:
		return v, nil
	case Readable:
		return v.MakeReader(), nil
	default:
		i, err := MakeIterator(vm, v)
		if err == nil {
			return newIterReader(vm, i), nil
		}
		return nil, fmt.Errorf("type %q %w", TypeName(v), ErrIsNotReadable)
	}
}

type Reader interface {
	Value
	io.Reader
}

type emptyReaderImpl struct{}

var emptyReader = &emptyReaderImpl{}

func (r *emptyReaderImpl) String() string { return "<reader>" }
func (r *emptyReaderImpl) Type() string   { return "reader" }
func (r *emptyReaderImpl) Traits() Traits { return TraitNone }

func (r *emptyReaderImpl) MakeReader() (Reader, error) {
	return r, nil
}

func (r *emptyReaderImpl) Read(b []byte) (int, error) {
	return 0, io.EOF
}

type iterReader struct {
	m   *VM
	e   Iterator
	buf bytequeue.ByteQueue
}

func newIterReader(vm *VM, iter Iterator) *iterReader {
	return &iterReader{
		m: vm,
		e: iter,
	}
}

func (r *iterReader) String() string { return "<reader>" }
func (r *iterReader) Type() string   { return "reader" }
func (r *iterReader) Traits() Traits { return TraitNone }

func (r *iterReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	for len(r.buf.Peek()) < len(b) {
		v, err := r.m.IterNext(r.e, 1)
		if err != nil {
			return 0, err
		}
		if v == nil {
			break
		}
		r.buf.Write([]byte(v[0].String()))
		r.buf.Write([]byte{'\n'})
	}

	if len(r.buf.Peek()) == 0 {
		return 0, io.EOF
	}

	n := len(r.buf.Peek())
	if n > len(b) {
		n = len(b)
	}

	copy(b, r.buf.Peek()[:n])
	r.buf.Pop(n)

	return n, nil
}
