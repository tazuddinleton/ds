package singlyoo

import (
	"errors"
	"reflect"
)

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
	list.InsertAfter(list.Last(), data)
}

func (list *LinkedList[T]) Last() *Node[T] {
	h := list.head
	for h.next != nil {
		h = h.next
	}
	return h
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

// Pop removes first node from the list and returns it
func (list *LinkedList[T]) Pop() *Node[T] {
	n := list.head
	list.head = n.next
	return n
}

// Delete removes first matching element form the list
func (list *LinkedList[T]) Delete(data T) {
	if reflect.DeepEqual(list.head.data, data) {
		list.head = list.head.next
		return
	}
	prev := list.head
	next := list.head.next

	for !reflect.DeepEqual(next.data, data) && next != nil {
		prev = next
		next = next.next
	}

	if next != nil {
		prev.next = next.next
	}
}

// DeleteWhile removes nodes from the head while matches a condition
func (list *LinkedList[T]) DeleteWhile(selector func(*Node[T]) bool) {
	h := list.head

	for selector(h) {
		h = h.next
	}
	list.head = h
}

// DeleteWhen removes all those nodes when matches a condition
func (list *LinkedList[T]) DeleteWhen(selector func(*Node[T]) bool) {
	prev := list.head
	next := prev.next
	for next != nil {
		if selector(next) {
			prev.next = next.next
		} else {
			prev = next
		}
		next = prev.next
	}
	if selector(list.head) {
		list.head = list.head.next
	}
}

// DeleteAt takes an index and remove item at that index
func (list *LinkedList[T]) DeleteAt(index int) {
	if index == 0 {
		list.head = list.head.next
		return
	}
	cnt := 1
	prev := list.head
	next := prev.next
	for cnt < index {
		prev = next
		next = prev.next
		cnt++
	}
	prev.next = next.next
}

// Item takes an index at returns node at that index
func (list *LinkedList[T]) Item(index int) (T, error) {
	var t T
	if index < 0 {
		return t, errors.New("index out of range")
	}
	n := list.head
	cnt := 0
	for cnt < index {
		cnt++
		if n == nil {
			return t, errors.New("index out of range")
		}
		n = n.next
	}
	return n.data, nil
}

// 1, 2, 3, 4
// 1-> nil
// ReverseIter reverses the elements of a list
func (list *LinkedList[T]) ReverseIter() {
	var prev *Node[T]
	curr := list.head

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	list.head = prev
}

// 1,2,3,4
func (list *LinkedList[T]) ReverseIterFunc() {
	var reverse func(prev, curr *Node[T]) *Node[T]
	reverse = func(prev, curr *Node[T]) *Node[T] {
		if curr == nil {
			return prev
		}
		n := curr.next
		curr.next = prev
		return reverse(curr, n)
	}
	list.head = reverse(nil, list.head)
}

// 1,2,3,4,5
func (list *LinkedList[T]) ReverseRecurr() {
	var reverse func(head *Node[T]) *Node[T]
	reverse = func(head *Node[T]) *Node[T] {
		newHead := head
		if head.next != nil {
			newHead = reverse(head.next)
			head.next.next = head
		}
		head.next = nil
		return newHead
	}
	list.head = reverse(list.head)
}
