package fiber

type fiberListElem struct {
	Self       *Fiber
	Prev, Next *fiberListElem
}

func (n *fiberListElem) InsertAfter(e *fiberListElem) {
	n.Next = e.Next
	e.Next.Prev = n
	e.Next = n
	n.Prev = e
}

func (n *fiberListElem) InsertBefore(e *fiberListElem) {
	n.Next = e
	n.Prev = e.Prev
	e.Prev.Next = n
	e.Prev = n
}

func (e *fiberListElem) Remove() {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Prev = nil
	e.Next = nil
}

type fiberList struct {
	l fiberListElem
}

func (l *fiberList) Init() {
	l.l.Prev = l.End()
	l.l.Next = l.End()
}

func (l *fiberList) Begin() *fiberListElem {
	return l.l.Next
}

func (l *fiberList) End() *fiberListElem {
	return &l.l
}

func (l *fiberList) Empty() bool {
	return l.Begin() == l.End()
}
