package doubly

type List[T any] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
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
	return &List[T]{Head: head, Tail: tail}
}

// Find takes a selector func and returns the first match using that func or nil
func (l *List[T]) Find(match func(item T) bool) *ListNode[T] {
	h := l.Head
	for h != nil && !match(h.Val) {
		h = h.Next
	}
	return h
}
