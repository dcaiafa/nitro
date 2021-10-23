package ioqueue

import (
	"sync"
	"sync/atomic"
)

type Data struct {
	Data []byte
	Err  error
	arr  [512]byte
}

var ioRequestPool = sync.Pool{
	New: func() interface{} {
		return &Data{}
	},
}

func NewData() *Data {
	r := ioRequestPool.Get().(*Data)
	r.Data = r.arr[:]
	return r
}

type IOQueue struct {
	C      chan *Data
	closer sync.Once
	closed int32
}

func New() *IOQueue {
	return &IOQueue{
		C: make(chan *Data, 1),
	}
}

func (q *IOQueue) Close() {
	q.closer.Do(func() {
		close(q.C)
		atomic.StoreInt32(&q.closed, 1)
	})
}

func (q *IOQueue) Closed() bool {
	return atomic.LoadInt32(&q.closed) != 0
}

func ReleaseData(r *Data) {
	ioRequestPool.Put(r)
}
