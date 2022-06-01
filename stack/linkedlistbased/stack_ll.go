package stackll

import "errors"

type Stack[T any] struct {
	store *node[T]
}
type node[T any] struct {
	data T
	next *node[T]
}

func NewStack[T any](len int) *Stack[T] {
	return &Stack[T]{}
}

// Empty returns true if stack is empty, false otherwise
// Time = O(1), Space = O(1)
func (s *Stack[T]) Empty() bool {
	return s.store == nil
}

// Push pushes an item at the top of the stack
// Time = O(1), Space = O(1)
func (s *Stack[T]) Push(item T) {
	s.store = &node[T]{data: item, next: s.store}
}

// Pop returns a top item from the stack, or error if stack is empty
// Time = O(1), Space = O(1)
func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var t T
		return t, errors.New("stack is empty")
	}
	t := s.store.data
	s.store = s.store.next
	return t, nil
}
