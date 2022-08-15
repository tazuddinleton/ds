package doubly

import "testing"

func TestListOf(t *testing.T) {
	l := ListOf(1, 2, 3)

	if l.head.Val != 1 {
		t.Errorf("wanted 1, got %d", l.head.Val)
	}

	if l.Tail().Val != 3 {
		t.Errorf("wanted 3, got %d", l.Tail().Val)
	}
}

func TestFind(t *testing.T) {
	l := ListOf(1, 2, 3, 4, 5, 6, 7)
	node := l.Find(func(item int) bool { return item > 5 && item < 10 })
	if node.Val != 6 {
		t.Errorf("wanted 6, got %d", node.Val)
	}
}

func TestInsert(t *testing.T) {
	l := ListOf(1)
	l.Insert(4)
	if l.head.Val != 4 {
		t.Errorf("wanted 4, got %d", l.head.Val)
	}

	l2 := &List[int]{head: nil}

	l2.Insert(1)
	if l2.Head().Val != 1 {
		t.Errorf("wanted 1, got %d", l2.Head().Val)
	}
	if l2.Tail().Val != 1 {
		t.Errorf("wanted 1, got %d", l2.Tail().Val)
	}
}
