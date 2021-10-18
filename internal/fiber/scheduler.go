package fiber

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type fiberState int

const (
	stateNew fiberState = iota
	stateRunning
	stateReady
	stateBlocking
	stateTerminated
)

type Fiber struct {
	Data interface{}

	le    fiberListElem
	id    uint32
	sched *Scheduler
	cv    *sync.Cond
	state fiberState
	fn    func()

	blockCancel context.CancelFunc
}

func (f *Fiber) expectState(state fiberState) {
	if f.state != state {
		panic(fmt.Sprintf("Expected state %v but it was %v", state, f.state))
	}
}

func newFiber(sched *Scheduler, id uint32, fn func()) *Fiber {
	f := &Fiber{
		id:    id,
		sched: sched,
		state: stateNew,
		fn:    fn,
	}
	f.cv = sync.NewCond(&sched.mutex)
	f.le.Self = f
	return f
}

type Scheduler struct {
	mutex   sync.Mutex
	cv      *sync.Cond
	ready   fiberList
	blocked fiberList
	active  *Fiber
	lastID  uint32
}

func NewScheduler() *Scheduler {
	s := &Scheduler{}
	s.cv = sync.NewCond(&s.mutex)
	s.ready.Init()
	s.blocked.Init()
	return s
}

func (s *Scheduler) Run(main *Fiber) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	main.state = stateRunning
	s.active = main
	s.runFiber(main)

	for {
		// Wait while a fiber is currently running or while there is nothing ready.
		for s.active != nil || (s.ready.Empty() && !s.blocked.Empty()) {
			s.cv.Wait()
		}

		if s.ready.Empty() && s.blocked.Empty() {
			break
		}

		f := s.ready.Begin().Self
		f.le.Remove()

		s.active = f
		f.state = stateRunning

		s.mutex.Unlock()
		f.cv.Signal()
		s.mutex.Lock()
	}
}

func (s *Scheduler) Block(ctx context.Context, fn func(ctx context.Context)) {
	s.mutex.Lock()

	me := s.active

	ctx, me.blockCancel = context.WithCancel(ctx)
	defer me.blockCancel()

	s.active = nil
	me.state = stateBlocking
	me.le.InsertBefore(s.blocked.End())

	s.mutex.Unlock()
	s.cv.Signal()

	fn(ctx)

	s.mutex.Lock()
	me.state = stateReady
	me.le.Remove()
	me.le.InsertBefore(s.ready.End())
	s.mutex.Unlock()
	s.cv.Signal()

	s.mutex.Lock()
	for me.state != stateRunning {
		me.cv.Wait()
	}
	s.mutex.Unlock()

	me.blockCancel = nil
}

func (s *Scheduler) NewFiber(fn func()) *Fiber {
	id := atomic.AddUint32(&s.lastID, 1)
	return newFiber(s, id, fn)
}

func (s *Scheduler) SwitchToNew(f *Fiber) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	f.expectState(stateNew)

	parked := s.active
	parked.state = stateReady
	parked.le.InsertBefore(s.ready.End())

	s.active = f
	f.state = stateRunning

	s.runFiber(f)

	for parked.state != stateRunning {
		parked.cv.Wait()
	}
}

func (s *Scheduler) ForEachFiber(h func(fiber *Fiber)) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.active != nil {
		h(s.active)
	}
	for le := s.ready.Begin(); le != s.ready.End(); le = le.Next {
		h(le.Self)
	}
	for le := s.blocked.Begin(); le != s.blocked.End(); le = le.Next {
		h(le.Self)
	}
}

func (s *Scheduler) CancelBlocked() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for le := s.blocked.Begin(); le != s.blocked.End(); le = le.Next {
		le.Self.blockCancel()
	}
}

func (s *Scheduler) runFiber(f *Fiber) {
	go func() {
		f.fn()

		s.mutex.Lock()
		s.active = nil
		f.state = stateTerminated
		s.mutex.Unlock()
		s.cv.Signal()
	}()
}
