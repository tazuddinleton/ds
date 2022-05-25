package singlyfp

import "fmt"

// Node describes a singly linked list data structure
type Node[T any] struct {
	data T
	next *Node[T]
}

// New creates a new singly linked list
func New[T any](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data, next}
}

// NewList a variadic constructor for singly linked list
func NewList[T any](ts ...T) *Node[T] {
	if len(ts) <= 0 {
		return nil
	}
	return New(ts[0], NewList((ts[1:])...))
}

// PrintList prints the data to console sequentially
func PrintList[T any](node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Print(node.data)
	PrintList(node.next)
}

// Returns a flat slice representation of the singly linked list
func Flatten[T any](node *Node[T]) []T {
	var loop func(*Node[T], []T) []T
	loop = func(n *Node[T], t []T) []T {
		if n == nil {
			return t
		}
		return loop(n.next, append(t, n.data))
	}
	return loop(node, []T{})
}

// FlattenSelect works similart to Flatten but lets user pass a selector to what data should be selected
func FlattenSelect[T any, P any](node *Node[T], f func(T) P) []P {
	var loop func(*Node[T], []P) []P
	loop = func(n *Node[T], t []P) []P {
		if n == nil {
			return t
		}
		return loop(n.next, append(t, f(n.data)))
	}
	return loop(node, []P{})
}

// InsertHead inserts a node at the beginning of the list
func InsertHead[T any](data T, node *Node[T]) *Node[T] {
	return New(data, node)
}

// Append appends a data node at the end of a given list
func Append[T any](list *Node[T], data T) *Node[T] {
	if list == nil {
		return New(data, nil)
	}
	return New(list.data, Append(list.next, data))
}

// Last returns the last node of a given list
func Last[T any](list *Node[T]) T {
	if list == nil {
		var t T
		return t
	}

	if list.next == nil {
		return list.data
	}
	return Last(list.next)
}

func InsertAfter[T any](data T, after T, node *Node[T]) *Node[T] {
	// this is not a valid operation
	return nil
}

func Reverse[T any](list *Node[T]) *Node[T] {
	if list == nil {
		return list
	}
	return Append(Reverse(list.next), list.data)
}
