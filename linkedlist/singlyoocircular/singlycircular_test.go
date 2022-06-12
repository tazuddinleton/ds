package singlyoocircular

import "testing"

func TestNewNode(t *testing.T) {
	l := newNode(1, newNode(2, nil))
	if l.data != 1 {
		t.Errorf("wanted 1 got %d", l.data)
	}
}

func TestNewList(t *testing.T) {
	h, tail := newList("Hello", "World", "!")
	if h.data != "Hello" {
		t.Errorf("wanted \"Hello\", got %s", h.data)
	}
	if tail.data != "!" {
		t.Errorf("wanted !, got %s", tail.data)
	}
}

func TestNew_List(t *testing.T) {
	l := NewList[int]()

	if l.head != nil {
		t.Errorf("wanted nil, got %v", l.head)
	}
	if l.tail != nil {
		t.Errorf("wanted nil, got %v", l.head)
	}
	l.Push(1)
	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}
	if l.tail.data != 1 {
		t.Errorf("wanted 1, got %d", l.tail.data)
	}
	if l.Count() != 1 {
		t.Errorf("wanted 1, got %d", l.Count())
	}

	l.Push(2)
	if l.Count() != 2 {
		t.Errorf("wanted 2, got %d", l.Count())
	}
	if l.head.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.data)
	}
}

func TestNew_ListOf(t *testing.T) {
	l := NewListOf(1, 2, 3, 4)
	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}

	if l.Count() != 4 {
		t.Errorf("wanted 4, got %d", l.Count())
	}

	if l.tail.data != 4 {
		t.Errorf("wanted 4, got %d", l.tail.data)
	}
}

func TestPush(t *testing.T) {
	l := NewListOf(1, 2, 3)
	if l.Count() != 3 {
		t.Errorf("wanted 3, got %d", l.Count())
	}
	l.Push(10)

	if l.head.data != 10 {
		t.Errorf("wanted 10, got %d", l.head.data)
	}

	if l.Count() != 4 {
		t.Errorf("wanted 4, got %d", l.Count())
	}
}

func TestPushAfter(t *testing.T) {
	l := NewListOf(1, 2, 3)

	two := l.head.next

	l.InsertAfter(two, 4)

	if two.next.data != 4 {
		t.Errorf("wanted 4, got %d", two.next.data)
	}

	if l.Count() != 4 {
		t.Errorf("wanted 4, got %d", l.Count())
	}

	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}

	if l.tail.data != 3 {
		t.Errorf("wanted 3, got %d", l.tail.data)
	}

	l.InsertAfter(l.tail, 100)

	if l.head.data != 1 {
		t.Errorf("wanted 1, got %d", l.head.data)
	}

	if l.tail.data != 100 {
		t.Errorf("wanted 100, got %d", l.tail.data)
	}
}

func TestAppend(t *testing.T) {
	l := NewListOf(1)
	l.Append(2)
	if l.tail.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.next.data)
	}
	l.Append(3)
	if l.tail.data != 3 {
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

func TestLast(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5)
	lst := l.Last()
	if lst.data != 5 {
		t.Errorf("wanted 5, got %d", lst.data)
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

func TestDeleteWhile(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5)
	l.DeleteWhile(func(n *Node[int]) bool { return n.data < 3 }) // 3,4,5

	if l.head.data != 3 {
		t.Errorf("wanted 3, got %d", l.head.data)
	}
}

func TestDeleteWhen(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 4, 4, 4, 5, 6)
	l.DeleteWhen(func(n *Node[int]) bool { return n.data == 3 || n.data == 4 }) // 1, 2, 5, 6
	if l.head.next.next.data != 5 {
		t.Errorf("wanted 5, got %d", l.head.next.next.data)
	}

	l.DeleteWhen(func(n *Node[int]) bool { return n.data == 1 })

	if l.head.data != 2 {
		t.Errorf("wanted 2, got %d", l.head.data)
	}
	if l.Count() != 3 {
		t.Errorf("wanted 3, got %d", l.Count())
	}
}

func TestItem(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5, 6)
	one, _ := l.Item(0)
	if one != 1 {
		t.Errorf("wanted 1, got %d", one)
	}

	six, _ := l.Item(5)
	if six != 6 {
		t.Errorf("wanted 6, got %d", six)
	}

	_, err := l.Item(-10)
	if err == nil {
		t.Errorf("wanted index out of range, got nil")
	}
	_, err = l.Item(1000)
	if err == nil {
		t.Errorf("wanted index out of range, got nil")
	}
}

func TestDeleteAt(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5, 6)
	l.DeleteAt(2) // 1, 2, 4, 5, 6
	i, _ := l.Item(2)
	if i != 4 {
		t.Errorf("wanted 4, got %d", i)
	}
}

func TestString(t *testing.T) {
	l := NewListOf(1, 2, 3, 4, 5, 6)

	str := l.String()

	if str != "123456" {
		t.Errorf("wanted 123456, got %s", str)
	}
}
