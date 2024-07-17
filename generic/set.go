package generic

/*

  File:    set.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: generic set implementation. Index is insertion order (1 relative).
*/

type Set[T comparable] struct {
	items []T
}

// NewSet of T
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{nil}
}

// Contains - returns index of an existing T.  -1 if not found
func (s *Set[T]) Contains(t T) int {
	for ix := 0; ix < len(s.items); ix++ {
		if s.items[ix] == t {
			return ix
		}
	}
	return -1
}

// Add a new unique T. old index if already exists
func (s *Set[T]) Add(t T) int {
	ix := s.Contains(t)
	if ix == -1 {
		s.items = append(s.items, t)
		ix = len(s.items) - 1
	}
	return ix
}

// Remove the T at position ix
func (s *Set[T]) Remove(ix int) (b bool) {
	if ix > -1 && ix < len(s.items) {
		s.items = append(s.items[:ix], s.items[ix+1:]...)
		b = true
	}
	return
}

// Get the T at index position ix
func (s *Set[T]) Get(ix int) (t T, b bool) {
	ix--
	if ix > -1 && ix < len(s.items) {
		t = s.items[ix]
		b = true
	}
	return
}

// Count (size) of the Set
func (s *Set[T]) Count() int {
	return len(s.items)
}

// Clear - remove all T items on the stack
func (s *Set[T]) Clear() {
	s.items = nil
}
