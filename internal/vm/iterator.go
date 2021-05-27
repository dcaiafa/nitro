package vm

import "fmt"

type Iterator interface {
	Callable
	isIterator()
	IterNRet() int
	Close() error
}

type ILIterator struct {
	fn         *Fn
	captures   []ValueRef
	iterNRet   int
	tryCatches []tryCatch
	defers     []*Closure
	stack      []Value
	nlocals    int
	ip         int
	sp         int

	preAllocStack [stackSize]Value
}

var _ Iterator = (*ILIterator)(nil)

func (e *ILIterator) isIterator()    {}
func (e *ILIterator) String() string { return "<Iterator>" }
func (e *ILIterator) Type() string   { return "Iterator" }

func (e *ILIterator) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (i *ILIterator) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Iterator calls are handled directly by the VM.
	panic("not called")
}

func (i *ILIterator) IterNRet() int { return i.iterNRet }

func (i *ILIterator) Close() error { return nil }

type NativeIterator struct {
	extFn    NativeFn
	closeFn  NativeFn
	iterNRet int
}

var _ Iterator = (*NativeIterator)(nil)

func (e *NativeIterator) isIterator()    {}
func (e *NativeIterator) String() string { return "<Iterator>" }
func (e *NativeIterator) Type() string   { return "Iterator" }

func (e *NativeIterator) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (i *NativeIterator) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Iterator calls are handled directly by the VM.
	panic("not called")
}

func (i *NativeIterator) IterNRet() int { return i.iterNRet }

func (i *NativeIterator) Close() error { return nil }

func NewIterator(extFn, closeFn NativeFn, nret int) Iterator {
	i := &NativeIterator{
		extFn:    extFn,
		closeFn:  closeFn,
		iterNRet: nret,
	}
	return i
}
