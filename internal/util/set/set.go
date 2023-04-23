package set

type Set[T comparable] struct {
	set map[T]bool
}

func (s *Set[T]) Add(v T) {
	if s.set == nil {
		s.set = make(map[T]bool)
	}
	s.set[v] = true
}

func (s *Set[T]) Has(v T) bool {
	if s.set == nil {
		return false
	}
	return s.set[v]
}

func (s *Set[T]) ToSlice() []T {
	r := make([]T, 0, len(s.set))
	for v := range s.set {
		r = append(r, v)
	}
	return r
}
