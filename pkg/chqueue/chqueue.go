// implements a queue or FIFO list using buffered channels
package chqueue

import (
	"errors"
)

type Queue struct {
	Items chan int
}

func InitQueue(size int) *Queue {
	p := Queue{Items: make(chan int, size)}
	return &p
}

func Enqueue(queue *Queue, item int) error {
	select {
	case queue.Items <- item:
		return nil
	default:
		return errors.New("overflow")
	}
}

func Dequeue(queue *Queue) (item int, err error) {
	select {
	case result := <-queue.Items:
		return result, nil
	default:
		return -1, errors.New("underflow")
	}
}

func EmptyQueue(queue *Queue) bool {
	if len(queue.Items) == 0 {
		return true
	} else {
		return false
	}
}
