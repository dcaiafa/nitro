package vm

import (
	"context"

	"github.com/dcaiafa/nitro/internal/fiber"
)

type coroutine struct {
	fiber         *fiber.Fiber
	callStack     []*frame
	frame         *frame
	pkg           *CompiledPackage
	stack         []Value
	globals       []Value
	instrs        []Instr
	sp            int
	ip            int
	preAllocStack [stackSize]Value
	framePool     []*frame
	pendingErr    error
	contextStack  []context.Context
}

func (co *coroutine) TopContext() context.Context {
	return co.contextStack[len(co.contextStack)-1]
}

func (co *coroutine) PushFrame(frame *frame) {
	if len(co.callStack) > 0 {
		co.frame.ip = co.ip
	}
	co.callStack = append(co.callStack, frame)
	co.frame = frame
	if co.frame.fn != nil {
		co.pkg = co.frame.fn.pkg
		co.globals = co.pkg.globals
		co.instrs = co.frame.fn.instrs
		co.ip = co.frame.ip
	} else {
		co.instrs = nil
	}
}

func (co *coroutine) PopFrame() {
	f := co.callStack[len(co.callStack)-1]
	*f = frame{}
	co.framePool = append(co.framePool, f)

	co.callStack[len(co.callStack)-1] = nil
	co.callStack = co.callStack[:len(co.callStack)-1]

	if len(co.callStack) > 0 {
		co.frame = co.callStack[len(co.callStack)-1]
		if co.frame.fn != nil {
			co.pkg = co.frame.fn.pkg
			co.globals = co.pkg.globals
			co.instrs = co.frame.fn.instrs
			co.ip = co.frame.ip
		}
	}
}

func (co *coroutine) NewFrame() *frame {
	if len(co.framePool) == 0 {
		return &frame{}
	}
	f := co.framePool[len(co.framePool)-1]
	co.framePool = co.framePool[:len(co.framePool)-1]
	return f
}

func (co *coroutine) PushContext(ctx context.Context) {
	co.contextStack = append(co.contextStack, ctx)
}

func (co *coroutine) PopContext() context.Context {
	n := len(co.contextStack) - 1
	popped := co.contextStack[n]
	co.contextStack[n] = nil
	co.contextStack = co.contextStack[:n]
	return popped
}
