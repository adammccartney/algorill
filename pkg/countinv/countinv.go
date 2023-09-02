package countinv

import (
	"errors"

	"github.com/adammccartney/algorill/pkg/queue"
)

// Based on two way merge algorithm from Knuth
// use of queues inspired by Skiena
func MergeCountSplitInv(c []int, d []int) ([]int, int, error) {

	bufx := queue.InitQueue(len(c))
	bufy := queue.InitQueue(len(d))

	for _, val := range c {
		if err := bufx.Enqueue(val); err != nil {
			return nil, -1, err
		}
	}
	for _, val := range d {
		if err := bufy.Enqueue(val); err != nil {
			return nil, -1, err
		}
	}

	// M1: initialize
	var result []int
	splitinv := 0

	for !bufx.EmptyQueue() && !bufy.EmptyQueue() {
		// M2: find smaller
		// peak the head of each queue to compare
		if bufx.HeadQueue() <= bufy.HeadQueue() {
			// M3: output x
			x, err := bufx.Dequeue()
			if err != nil {
				return nil, -1, err
			}
			result = append(result, x)
		} else {
			y, err := bufy.Dequeue()
			if err != nil {
				return nil, -1, err
			}
			result = append(result, y)
			splitinv = splitinv + (len(bufy.Items) - len(bufx.Items) + 1)
		}
	}
	for !bufy.EmptyQueue() {
		y, err := bufy.Dequeue()
		if err != nil {
			return nil, -1, err
		}
		result = append(result, y)
	}
	for !bufx.EmptyQueue() {
		x, err := bufx.Dequeue()
		if err != nil {
			return nil, -1, err
		}
		result = append(result, x)
	}
	return result, splitinv, nil
}

// SortCountInv takes as input an array a of n integers.
// It counts the number of inversions in the array and returns the sorted array
// B, along with the number of inversions.
func SortCountInv(a []int) ([]int, int, error) {
	if len(a) <= 1 {
		return a, 0, nil
	} else {
		mid := len(a) / 2
		c, leftInv, err := SortCountInv(a[:mid])
		if err != nil {
			return []int{}, -1, errors.New("error counting inv left half")
		}
		d, rightInv, err := SortCountInv(a[mid:])
		if err != nil {
			return []int{}, -1, errors.New("error counting inv right half")
		}
		b, splitInv, err := MergeCountSplitInv(c, d)
		if err != nil {
			return []int{}, -1, errors.New("error merge counting split inv")
		}
		return b, leftInv + rightInv + splitInv, nil
	}
}
