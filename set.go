package aoc

type Set[T comparable] map[T]bool

func (s Set[T]) Add(item T) {
	s[item] = true
}
