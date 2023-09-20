package minstack

import (
	"errors"
)

// MinStack guarantees that the Minimum value is always at the top of the stack.
type MinStack struct {
	vals []int
	min  int
}

func NewMinStack() *MinStack {
	m := &MinStack{
		vals: make([]int, 0),
		min:  0,
	}
	return m
}

func (m *MinStack) Min() (int, error) {
	if len(m.vals) == 0 {
		return -1, errors.New("stack empty")
	} else {
		return m.min, nil
	}
}

func (m *MinStack) Pop() (int, error) {
	if len(m.vals) == 0 {
		return -1, errors.New("stack empty")
	}

	top := m.vals[len(m.vals)-1]
	m.vals = m.vals[:len(m.vals)-1] // pop

	if top < m.min { // new Min!
		result := m.min
		m.min = 2*m.min - top
		return result, nil
	} else {
		return top, nil
	}
}

func (m *MinStack) Push(n int) error {
	if len(m.vals) == 0 { // stack empty
		m.min = n
		m.vals = append(m.vals, n)
	} else if n < m.min { // new min!
		m.vals = append(m.vals, 2*n-m.min)
		m.min = n
	} else {
		m.vals = append(m.vals, n)
	}
	return nil
}
