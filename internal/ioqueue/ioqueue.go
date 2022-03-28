package ioqueue

import (
	"sync"
	"sync/atomic"
)

const (
	ioQueueClosed     uint32 = 0b01
	ioQueueWantsClose uint32 = 0b10
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
	flags  uint32
}

func New() *IOQueue {
	return &IOQueue{
		C: make(chan *Data, 1),
	}
}

func (q *IOQueue) Close() {
	q.closer.Do(func() {
		close(q.C)
		q.setFlag(ioQueueClosed)
	})
}

func (q *IOQueue) SignalWantsClose() {
	q.setFlag(ioQueueWantsClose)
}

func (q *IOQueue) setFlag(flag uint32) {
	for {
		oldFlags := atomic.LoadUint32(&q.flags)
		newFlags := oldFlags | flag
		if atomic.CompareAndSwapUint32(&q.flags, oldFlags, newFlags) {
			break
		}
	}
}

func (q *IOQueue) State() (closed, wantsClose bool) {
	flags := atomic.LoadUint32(&q.flags)
	closed = (flags & ioQueueClosed) != 0
	wantsClose = (flags & ioQueueWantsClose) != 0
	return closed, wantsClose
}

func ReleaseData(r *Data) {
	ioRequestPool.Put(r)
}
