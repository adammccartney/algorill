package mergesort

import (
	"errors"
	"fmt"

	"github.com/adammccartney/algorill/pkg/datastruct/queue"
)

// Two way merge algorithm, based on Knuth's version [taocp vol3 pg 158]
// sorts two sorted lists into one
func Merge(c []int, d []int) ([]int, error) {
	// declare two buffered channels to use
	bufx := queue.InitQueue(len(c))
	bufy := queue.InitQueue(len(d))

	for _, val := range c {
		if err := bufx.Enqueue(val); err != nil {
			return nil, err
		}
	}
	for _, val := range d {
		if err := bufy.Enqueue(val); err != nil {
			return nil, err
		}
	}

	// M1: Intialize
	// declare output array
	var result []int

	for !bufx.EmptyQueue() && !bufy.EmptyQueue() {
		// M2: Find smaller,
		// peak the head of each queue to compare
		if bufx.HeadQueue() <= bufy.HeadQueue() {
			// M3: Output x
			x, err := bufx.Dequeue()
			if err != nil {
				return nil, err
			}
			result = append(result, x)
		} else {
			// M5: Output y
			y, err := bufy.Dequeue()
			if err != nil {
				return nil, err
			}
			result = append(result, y)
		}
	}
	for !bufy.EmptyQueue() { // M4: transmit remaining y
		y, _ := bufy.Dequeue()
		result = append(result, y)
	}
	for !bufx.EmptyQueue() { // M6: transmit remaining x
		x, _ := bufx.Dequeue()
		result = append(result, x)
	}

	return result, nil
}

// Divide an array into to equally sized arrays
// input array a of len n
// output: array b, c of len n/2
func divide(a []int) (b []int, c []int, err error) {
	n := len(a)
	if n%2 != 0 {
		return []int{}, []int{}, errors.New("array length is not even")
	}
	mid := n / 2
	b = a[0:mid]
	c = a[mid:]
	return b, c, nil
}

// sort array and return
func Mergesort(a []int) (b []int, err error) {
	// base case
	if len(a) <= 1 {
		return a, nil
	}
	c, d, err := divide(a)
	if err != nil {
		return []int{}, errors.New("error running divide subroutine")
	}
	c, err = Mergesort(c)
	if err != nil {
		return []int{}, errors.New("error with first recursive call")
	}
	d, err = Mergesort(d)
	if err != nil {
		return []int{}, errors.New("error with second recursive call")
	}
	return Merge(c, d)
}

func main() {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6}
	a, err := Mergesort(a)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
