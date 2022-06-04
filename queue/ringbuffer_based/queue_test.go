package queuearr

import "testing"

func TestEmpty(t *testing.T) {
	q := NewQueue[int](10)
	if !q.Empty() {
		t.Errorf("wanted emtpy queue, got %v", q)
	}
}

func TestFull(t *testing.T) {
	q := NewQueue[int](5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	if !q.Full() {
		t.Errorf("wanted full, got empty")
	}
	q.Dequeue()
	q.Enqueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	if !q.Full() {
		t.Errorf("wanted full, got empty")
	}
}

func TestReset(t *testing.T) {
	q := NewQueue[int](5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Reset()
	if !q.Empty() {
		t.Errorf("wanted emtpy queue, got %v", q)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := NewQueue[int](10)
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("wanted error cannot dequeue an empty queue, got %v", err)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int](10)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.tail != 3 {
		t.Errorf("wanted tail = 3, got %d", q.tail)
	}
	for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}
	if q.tail != 3 {
		t.Errorf("wanted tail = 3, got %d", q.tail)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int](10)
	for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}
	_1, _ := q.Dequeue()
	if _1 != 1 {
		t.Errorf("wanted 1, got %d", _1)
	}
}
