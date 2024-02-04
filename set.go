package qforge

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(element T) {
	(*s)[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(*s, element)
}
