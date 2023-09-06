package queue

import (
	"testing"
)

func TestGoodQueue(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	q := InitQueue(len(nums))

	// enqueue the numbers
	for _, n := range nums {
		err := q.Enqueue(n)
		if err != nil {
			t.Errorf("error enqueueing %d: %v", n, err)
		}
	}

	// dequeue the numbers
	for _, n := range nums {
		item, err := q.Dequeue()
		if err != nil {
			t.Errorf("error dequeueing %d: %v", n, err)
		}
		if item != n {
			t.Errorf("expected %d, got %d", n, item)
		}
	}
}

func TestEnqueueOverflow(t *testing.T) {
	q := InitQueue(5)
	for i := 0; i < 5; i++ {
		err := q.Enqueue(i)
		if err != nil {
			t.Errorf("error enqueueing %d: %v", i, err)
		}
	}
	err := q.Enqueue(6)
	if err == nil {
		t.Error("expected overflow error")
	}
}

func TestDequeueUnderflow(t *testing.T) {
	q := InitQueue(5)
	_, err := q.Dequeue()
	if err == nil {
		t.Error("expected underflow error")
	}
}

func TestHeadQueue(t *testing.T) {
	q := InitQueue(5)
	for i := 0; i < 5; i++ {
		err := q.Enqueue(i)
		if err != nil {
			t.Errorf("error enqueueing %d: %v", i, err)
		}
	}
	if q.HeadQueue() != 0 {
		t.Errorf("expected 0, got %d", q.HeadQueue())
	}

	_, err := q.Dequeue()
	if err != nil {
		t.Errorf("error dequeueing: %v", err)
	}
	if q.HeadQueue() != 1 {
		t.Errorf("expected 1, got %d", q.HeadQueue())
	}
	_, err = q.Dequeue()
	if err != nil {
		t.Errorf("error dequeueing: %v", err)
	}
	if q.HeadQueue() != 2 {
		t.Errorf("expected 2, got %d", q.HeadQueue())
	}
}

func TestEmptyQueue(t *testing.T) {
	q := InitQueue(5)
	if !q.EmptyQueue() {
		t.Error("expected empty queue")
	}
	q.Enqueue(1)
	if q.EmptyQueue() {
		t.Error("expected non-empty queue")
	}
	q.Enqueue(2)
	q.Enqueue(3)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	if !q.EmptyQueue() {
		t.Error("expected empty queue")
	}
}
