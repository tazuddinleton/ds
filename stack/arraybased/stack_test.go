package stack

import "testing"

func TestEmpty(t *testing.T) {
	s := NewStack[int](10)
	if !s.Empty() {
		t.Errorf("wanted empty, got not empty")
	}
}

func TestPush(t *testing.T) {
	s := NewStack[int](2)
	if !s.Empty() {
		t.Errorf("wanted empty, got not empty")
	}
	s.Push(4)
	s.Push(5)
	if s.Empty() {
		t.Errorf("wanted not empty, got empty")
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int](2)
	_, err := s.Pop()
	if err == nil {
		t.Errorf("wanted stack is empty error, got %v", err)
	}
	s.Push(1)
	s.Push(2)

	_2, err := s.Pop()
	if err != nil {
		t.Errorf("wanted error nil, but got %v", err)
	}
	if _2 != 2 {
		t.Errorf("wanted 2, got %d", _2)
	}
	_1, _ := s.Pop()
	if _1 != 1 {
		t.Errorf("wanted 1, got %d", _1)
	}
	if !s.Empty() {
		t.Errorf("wanted empty stack, got non empty stack")
	}
}
