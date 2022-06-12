package singlyoocircular

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type LinkedList[T any] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
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
	h, t := newList(ds...)
	return &LinkedList[T]{head: h, tail: t, count: len(ds)}
}

// newList is a variadic constructor for singly linked list
func newList[T any](ds ...T) (h, t *Node[T]) {
	if len(ds) == 0 {
		return nil, nil
	}
	h = newNode(ds[0], nil)
	t = h
	for i := 1; i < len(ds); i++ {
		t.next = newNode(ds[i], nil)
		t = t.next
	}
	if reflect.DeepEqual(h, t) {
		return
	}
	t.next = h
	return
}

func (list *LinkedList[T]) Count() int {
	return list.count
}

// Push inserts a node at the beginning of the list
func (list *LinkedList[T]) Push(data T) {
	list.head = newNode(data, list.head)
	list.count++
	if list.count < 2 {
		list.tail = list.head
	} else {
		list.tail.next = list.head
	}
}

// InsertAfter inserts a node after a given node
func (list *LinkedList[T]) InsertAfter(node *Node[T], data T) {
	n := newNode(data, node.next)
	node.next = n
	list.count++
	if reflect.DeepEqual(list.tail, node) {
		list.tail = n
	}
}

// Append inserts a node at the end of the list
func (list *LinkedList[T]) Append(data T) {
	list.InsertAfter(list.tail, data)
}

func (list *LinkedList[T]) Last() *Node[T] {
	return list.tail
}

// newNode creates a new singly linked list
func newNode[T any](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data, next}
}

// Pop removes first node from the list and returns it
func (list *LinkedList[T]) Pop() *Node[T] {
	n := list.head
	list.head = n.next
	list.count--
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

	for !reflect.DeepEqual(next, list.head) {
		if reflect.DeepEqual(next.data, data) {
			prev.next = next.next
			next = prev.next.next
			list.count--
			break
		}
		prev = next
		next = next.next
	}
	if reflect.DeepEqual(next, list.head) {
		list.tail = prev
	}
}

// DeleteWhile removes nodes from the head while matches a condition
func (list *LinkedList[T]) DeleteWhile(selector func(*Node[T]) bool) {
	h := list.head

	for selector(h) {
		h = h.next
		list.count--
	}
	list.head = h
}

// DeleteWhen removes all those nodes when matches a condition
func (list *LinkedList[T]) DeleteWhen(selector func(*Node[T]) bool) {
	prev := list.head
	next := list.head.next
	for !reflect.DeepEqual(next, list.head) {
		if selector(next) {
			prev.next = next.next
			next = next.next
			list.count--
			continue
		}
		prev = next
		next = next.next
	}
	if reflect.DeepEqual(next, list.head) {
		list.tail = prev
	}

	if selector(list.head) {
		list.head = list.head.next
		list.count--
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
	if index < 0 || index >= list.count {
		return t, errors.New("index out of range")
	}
	n := list.head
	cnt := 0
	for cnt < index {
		cnt++
		n = n.next
	}
	return n.data, nil
}

func (list *LinkedList[T]) String() string {
	if list.head == nil {
		return ""
	}
	tmp := list.head
	str := new(strings.Builder)
	for {
		str.WriteString(fmt.Sprintf("%v", tmp.data))
		tmp = tmp.next
		if reflect.DeepEqual(tmp, list.head) {
			break
		}
	}
	return str.String()
}
