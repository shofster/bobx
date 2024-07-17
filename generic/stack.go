package generic

/*

  File:    stack.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*

   Description: generic stack implementation. LIFO

  for performance, new items are located @ end of slice

*/

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push a T onto top of Stack []T  (located @ bottom)
func (s *Stack[T]) Push(item T) int {
	s.items = append(s.items, item)
	return len(s.items)
}

// Peek (retained) the top T of Stack
func (s *Stack[T]) Peek() (item T, b bool) {
	if len(s.items) > 0 {
		item = s.items[len(s.items)-1]
		b = true
	}
	return
}

// Pop (removed) the top T of Stack.
func (s *Stack[T]) Pop() (item T, b bool) {
	if len(s.items) > 0 {
		item = s.items[len(s.items)-1]     // last item
		s.items = s.items[:len(s.items)-1] // remove it
		b = true
	}
	return
}

// Reverse iIterate through the Stack items. FIFO
func (s *Stack[T]) Reverse(iter func(item T)) {
	for _, v := range s.items { // skip index
		iter(v)
	}
}

// Count (size) of the Stack
func (s *Stack[T]) Count() int {
	return len(s.items)
}

// Clear - remove T items on the stack
func (s *Stack[T]) Clear() {
	s.items = nil
}
