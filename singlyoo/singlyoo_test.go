package singlyoo

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	l := newNode(1, newNode(2, nil))
	if l.data != 1 {
		t.Errorf("wanted 1 got %d", l.data)
	}
}

func TestNewList(t *testing.T) {
	l := newList("Hello", "World", "!")
	if l.data != "Hello" {
		t.Errorf("wanted \"Hello\", got %s", l.data)
	}
}

func TestNew_List(t *testing.T) {
	l := NewList[int]()

	if l.head != nil {
		t.Errorf("wanted nil, got %v", l.head)
	}
	l.Push(1)
	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}
}

func TestNew_ListOf(t *testing.T) {
	l := NewListOf(1, 2, 3, 4)
	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}
}

func TestPush(t *testing.T) {
	l := NewListOf(1, 2, 3)
	l.Push(10)

	if l.head.data != 10 {
		t.Errorf("wanted 10, got %d", l.head.data)
	}
}

func TestPushAfter(t *testing.T) {
	l := NewListOf(1, 2, 3)

	two := l.head.next

	l.InsertAfter(two, 4)

	if two.next.data != 4 {
		t.Errorf("wanted 4, got %d", two.next.data)
	}
}

func TestAppend(t *testing.T) {
	l := NewListOf(1)
	l.Append(2)
	if l.head.next.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.next.data)
	}
	l.Append(3)
	if l.head.next.next.data != 3 {
		t.Errorf("wanted 3, got %d", l.head.next.next.data)
	}
}

func TestPop(t *testing.T) {
	l := NewListOf(1, 2, 3)
	one := l.Pop()

	if l.head.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.data)
	}
	if one.data != 1 {
		t.Errorf("wanted 1, got %d", one.data)
	}
}

func TestDelete(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5)
	l.Delete(1) // 2,3,4,5
	if l.head.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.data)
	}

	l.Delete(5) // 2,3,4
	last := l.Last()
	if last.data != 4 {
		t.Errorf("wanted 4, got %d", last.data)
	}

	l.Delete(3) // // 2,4

	if l.head.next.data != 4 {
		t.Errorf("wanted 4, got %d", l.head.next.data)
	}
}
