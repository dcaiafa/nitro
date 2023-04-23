package stack

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}

func (s *Stack[T]) Empty() bool {
	return len(s.elems) == 0
}

func (s *Stack[T]) Push(e T) {
	s.elems = append(s.elems, e)
}

func (s *Stack[T]) Pop() T {
	l := len(s.elems)
	elem := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return elem
}
