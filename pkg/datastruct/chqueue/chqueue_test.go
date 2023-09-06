package chqueue

import (
	"fmt"
	"testing"
)

func TestGoodQueue(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	q := InitQueue(5)

	// enqueue 5 numbers
	for _, num := range nums {
		err := Enqueue(q, num)
		if err != nil {
			t.Errorf("Error enqueuing %d", num)
		}
	}

	// dequeue 5 numbers
	for i := 0; i < len(nums); i++ {
		a, err := Dequeue(q)
		if err != nil {
			fmt.Println(err)
			t.Errorf("Error dequeuing %d", i)
		}
		if a != nums[i] {
			t.Errorf("Expected %d, got %d", nums[i], a)
		}
	}
}

func TestEnqueueFullQueue(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	q := InitQueue(5)

	// enqueue 5 numbers
	for _, num := range nums {
		err := Enqueue(q, num)
		if err != nil {
			t.Errorf("Error enqueuing %d", num)
		}
	}

	// try to enqueue another number
	err := Enqueue(q, 6)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := InitQueue(5)

	// try to dequeue
	_, err := Dequeue(q)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
