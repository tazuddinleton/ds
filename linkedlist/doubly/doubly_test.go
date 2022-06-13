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

func TestFind(t *testing.T) {
	l := ListOf(1, 2, 3, 4, 5, 6, 7)
	node := l.Find(func(item int) bool { return item > 5 && item < 10 })
	if node.Val != 6 {
		t.Errorf("wanted 6, got %d", node.Val)
	}
}
