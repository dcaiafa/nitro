package vm

import (
	"fmt"
	"io"
)

type Readable interface {
	Value
	MakeReader() (Reader, error)
}

func IsRedable(v Value) bool {
	if v == nil {
		return true
	}
	if _, ok := v.(Readable); ok {
		return true
	}
	return false
}

func MakeReader(v Value) (Reader, error) {
	if v == nil {
		return emptyReader, nil
	}
	if r, ok := v.(Readable); ok {
		return r.MakeReader()
	}
	return nil, fmt.Errorf("value of type %q %w", TypeName(v), ErrIsNotReadable)
}

type Reader interface {
	Readable
	io.Reader
}

type emptyReaderImpl struct {
	BaseValue
}

var emptyReader = &emptyReaderImpl{
	BaseValue: BaseValue{TypeName: "reader"},
}

func (r *emptyReaderImpl) MakeReader() (Reader, error) {
	return r, nil
}

func (r *emptyReaderImpl) Read(b []byte) (int, error) {
	return 0, io.EOF
}
