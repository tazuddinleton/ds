package doubly

import "testing"

func TestListOf(t *testing.T) {
	l := ListOf(1, 2, 3)

	if l.Head.Val != 1 {
		t.Errorf("wanted 1, got %d", l.Head.Val)
	}

	if l.Tail.Val != 3 {
		t.Errorf("wanted 3, got %d", l.Tail.Val)
	}
}
