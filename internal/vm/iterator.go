package vm

import "fmt"

type Iterator interface {
	Callable
	IterNRet() int
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

type NativeIterator struct {
	extFn    NativeFn
	iterNRet int
}

var _ Iterator = (*NativeIterator)(nil)

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

func NewIterator(extFn NativeFn, nret int) Iterator {
	i := &NativeIterator{
		extFn:    extFn,
		iterNRet: nret,
	}
	return i
}
