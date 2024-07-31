package lib

import (
	"io"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

func take(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}
	if nitro.IsIterable(args[0]) {
		return takeIter(m, args, nRet)
	} else {
		return takeReader(m, args, nRet)
	}
}

func takeIter(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	inIter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	count, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	takeIter := &takeIterator{inIter: inIter, count: int(count)}
	return []nitro.Value{nitro.NewIterator(takeIter.Next, takeIter.Close, inIter.IterNRet())}, nil
}

type takeIterator struct {
	inIter nitro.Iterator
	count  int
}

func (i *takeIterator) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if i.count == 0 {
		i.Close(m)
		return nil, nil
	}

	v, err := m.IterNext(i.inIter, i.inIter.IterNRet())
	if err != nil {
		i.Close(m)
		return nil, err
	}
	if v == nil {
		i.Close(m)
		return nil, nil
	}
	i.count--
	return v, nil
}

func (i *takeIterator) Close(m *nitro.VM) error {
	m.IterClose(i.inIter)
	return nil
}

func takeReader(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	inReader, err := getReaderArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	count, err := getIntArg(args, 1)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{NewLimitedReader(inReader, count)}, nil
}

type LimitedReader struct {
	baseReader    vm.Reader
	limitedReader io.Reader
}

func NewLimitedReader(r vm.Reader, n int64) *LimitedReader {
	return &LimitedReader{
		baseReader:    r,
		limitedReader: io.LimitReader(r, n),
	}
}

func (r *LimitedReader) String() string    { return r.baseReader.String() }
func (r *LimitedReader) Type() string      { return r.baseReader.Type() }
func (r *LimitedReader) Traits() vm.Traits { return vm.TraitNone }

func (r *LimitedReader) Read(buf []byte) (int, error) {
	return r.limitedReader.Read(buf)
}

func (r *LimitedReader) Close() error {
	if closer, ok := r.baseReader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
