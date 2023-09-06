// simple FIFO queue implementation
package queue

import "errors"

type Queue struct {
	Items []int // slice of Items
	Front int   // index of the front of the queue
	Rear  int   // index of the rear of the queue
}

func InitQueue(size int) *Queue {
	return &Queue{
		Items: make([]int, size),
		Front: -1,
		Rear:  -1,
	}
}

func (q *Queue) Enqueue(item int) error {
	if q.Rear == len(q.Items)-1 {
		return errors.New("overflow")
	}
	q.Rear++
	q.Items[q.Rear] = item
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.Front == q.Rear {
		return -1, errors.New("underflow")
	}
	q.Front++
	return q.Items[q.Front], nil
}

func (q *Queue) EmptyQueue() bool {
	return q.Front == q.Rear
}

func (q *Queue) HeadQueue() int {
	return q.Items[q.Front+1]
}
