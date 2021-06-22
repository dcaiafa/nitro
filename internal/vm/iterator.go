package vm

import "fmt"

type Iterable interface {
	Value
	MakeIterator() Iterator
}

func MakeIterator(m *VM, v Value) (Iterator, error) {
	if v == nil {
		return NewIterator(emptyIter, nil, 1), nil
	}
	if i, ok := v.(Iterable); ok {
		return i.MakeIterator(), nil
	}
	return nil, fmt.Errorf("value of type %q %w", TypeName(v), ErrIsNotIterable)
}

type Iterator interface {
	Iterable
	IterNRet() int
	isIterator()
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

func (e *ILIterator) isIterator()            {}
func (e *ILIterator) String() string         { return "<iterator>" }
func (e *ILIterator) Type() string           { return "iterator" }
func (i *ILIterator) MakeIterator() Iterator { return i }

func (e *ILIterator) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (i *ILIterator) IterNRet() int { return i.iterNRet }

type CloseFn func(m *VM) error

type NativeIterator struct {
	extFn    NativeFn
	closeFn  CloseFn
	iterNRet int
	closed   bool
}

var _ Iterator = (*NativeIterator)(nil)

func (i *NativeIterator) isIterator()            {}
func (i *NativeIterator) String() string         { return "<Iterator>" }
func (i *NativeIterator) Type() string           { return "Iterator" }
func (i *NativeIterator) MakeIterator() Iterator { return i }

func (i *NativeIterator) EvalOp(op Op, operand Value) (Value, error) {
	return nil, fmt.Errorf("iterator does not support this operation")
}

func (i *NativeIterator) IterNRet() int { return i.iterNRet }
func (i *NativeIterator) Closed() bool  { return i.closed }

func (i *NativeIterator) Close(m *VM) error {
	if i.closed {
		return nil
	}
	i.closed = true
	if i.closeFn != nil {
		err := i.closeFn(m)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewIterator(extFn NativeFn, closeFn CloseFn, nret int) Iterator {
	i := &NativeIterator{
		extFn:    extFn,
		closeFn:  closeFn,
		iterNRet: nret,
	}
	return i
}

func emptyIter(m *VM, args []Value, nret int) ([]Value, error) {
	return nil, nil
}
