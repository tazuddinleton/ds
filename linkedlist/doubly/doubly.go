package doubly

type List[T any] struct {
	head *ListNode[T]
}

// ListNode[T any] defines a doubly linked list node
type ListNode[T any] struct {
	Val  T
	Prev *ListNode[T]
	Next *ListNode[T]
}

// ListOf[T any] is a variadic constructor of List
func ListOf[T any](nodes ...T) *List[T] {
	if len(nodes) == 0 {
		return nil
	}
	head := &ListNode[T]{Val: nodes[0]}
	tail := head
	for i := 1; i < len(nodes); i++ {
		tail.Next = &ListNode[T]{Val: nodes[i], Prev: tail}
		tail = tail.Next
	}
	return &List[T]{head: head}
}

func (l *List[T]) Head() *ListNode[T] {
	return l.head
}

func (l *List[T]) Tail() *ListNode[T] {
	if l.head == nil {
		return nil
	}
	return l.head.Prev
}

// Find takes a selector func and returns the first match using that func or nil
func (l *List[T]) Find(match func(item T) bool) *ListNode[T] {
	h := l.head
	for h != nil && !match(h.Val) {
		h = h.Next
	}
	return h
}

func (l *List[T]) Insert(item T) {
	h := l.head
	n := &ListNode[T]{Val: item, Next: h}
	l.head = n
	if h != nil {
		h.Prev = l.head
	}
}
