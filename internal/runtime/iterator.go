package runtime

type Iterator struct {
	fn         *Fn
	extFn      NativeFn
	captures   []ValueRef
	iterNRet   int
	locals     []Value
	tryCatches []tryCatch
	defers     []*Closure
	ip         int
}

func (e *Iterator) String() string { return "<Iterator>" }
func (e *Iterator) Type() string   { return "Iterator" }
func (e *Iterator) isCallable()    {}
func (e *Iterator) IterNRet() int  { return e.iterNRet }

func NewIterator(extFn NativeFn, caps []ValueRef, nret int) *Iterator {
	return &Iterator{
		extFn:    extFn,
		captures: caps,
		iterNRet: nret,
	}
}
