package minstack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := NewMinStack()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	stack.Push(4)

	if min, err := stack.Min(); err != nil || min != 1 {
		t.Errorf("Expected min of 1, got %d", min)
	}
	v, _ := stack.Pop()
	if v != 4 {
		t.Errorf("Expected 4, got %d", v)
	}
	v, _ = stack.Pop()
	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}

	if min, err := stack.Min(); err != nil || min != 2 {
		t.Errorf("Expected min of 2, got %d", min)
	}

}
