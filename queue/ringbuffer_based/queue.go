package queuearr

import "errors"

type Queue[T any] struct {
	store []T
	head  int
	tail  int
	full  bool
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{store: make([]T, size), head: 0, tail: 0, full: false}
}

func (q *Queue[T]) Empty() bool {
	return q.head == q.tail && !q.full
}

func (q *Queue[T]) Full() bool {
	return q.full
}

func (q *Queue[T]) Capacity() int {
	return len(q.store)
}

func (q *Queue[T]) Enqueue(item T) {
	q.store[q.tail] = item
	q.tail = (q.tail + 1) % len(q.store) // modulus is expensive computation
	if !q.full {
		q.full = q.tail == q.head
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	var t T
	if q.Empty() {
		return t, errors.New("cannot dequeue an empty queue")
	}
	q.full = false
	t = q.store[q.head]
	q.head = (q.head + 1) % len(q.store) // modulus is expensive computation
	if q.head == q.tail {
		q.Reset()
	}
	return t, nil
}

func (q *Queue[T]) Reset() {
	q.full = false
	q.head = 0
	q.tail = 0
}
