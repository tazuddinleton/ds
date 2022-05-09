package singlyoo

type LinkedList[T any] struct {
	head *Node[T]
}

// Node represents a singly linked list node
type Node[T any] struct {
	data T
	next *Node[T]
}

// NewList creates an empty linked list
func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// NewListOf creates a new linked list form the parameters
func NewListOf[T any](ds ...T) *LinkedList[T] {
	l := newList(ds...)
	return &LinkedList[T]{head: l}
}

// Push inserts a node at the beginning of the list
func (list *LinkedList[T]) Push(data T) {
	list.head = newNode(data, list.head)
}

// InsertAfter inserts a node after a given node
func (list *LinkedList[T]) InsertAfter(node *Node[T], data T) {
	n := newNode(data, node.next)
	node.next = n
}

// Append inserts a node at the end of the list
func (list *LinkedList[T]) Append(data T) {
	list.InsertAfter(list.last(list.head), data)
}

func (list *LinkedList[T]) last(l *Node[T]) *Node[T] {
	if l.next == nil {
		return l
	}
	return list.last(l.next)
}

// newNode creates a new singly linked list
func newNode[T any](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data, next}
}

// newList is a variadic constructor for singly linked list
func newList[T any](ds ...T) *Node[T] {
	if len(ds) == 0 {
		return nil
	}
	return newNode(ds[0], newList((ds[1:])...))
}
