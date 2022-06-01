package stack

import "errors"

type Stack[T any] struct {
	store []T
	top   int
}

func NewStack[T any](len int) *Stack[T] {
	return &Stack[T]{top: -1, store: make([]T, len)}
}

// Empty returns true if stack is empty, false otherwise
// Time = O(1), Space = O(1)
func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

// Push pushes an item at the top of the stack
// Time = O(1), Space = O(1)
func (s *Stack[T]) Push(item T) {
	s.top += 1
	s.store[s.top] = item
}

// Pop returns a top item from the stack, or error if stack is empty
// Time = O(1), Space = O(1)
func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var t T
		return t, errors.New("stack is empty")
	}
	s.top -= 1
	return s.store[s.top+1], nil
}
