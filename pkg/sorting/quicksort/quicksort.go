package quicksort

import (
	"errors"
	"log"
	"math/rand"
)

func Quicksort(A []int, l int, r int) {
	// Q1. initialize
	if l >= r { // 0- or 1-element subarray
		return
	}
	i, err := ChoosePivot(A, l, r)
	if err != nil {
		log.Fatal(err)
	}
	Swap(&A[l], &A[i]) // make pivot first

	j := Partition(A, l, r) // p = new pivot
	Quicksort(A, l, j-1)    // recurse on first part
	Quicksort(A, j+1, r)    // recurse on second part
}

func Swap(l *int, r *int) {
	if *l == *r {
		return
	}
	*r = *l ^ *r
	*l = *l ^ *r
	*r = *l ^ *r
}

func ChoosePivot(A []int, l int, r int) (int, error) {
	if l < 0 || r >= len(A) || l > r {
		return -1, errors.New("invalid bounds for pivot selection")
	}
	rnum := rand.Intn(r - l) // chose a uniform random number
	return l + rnum, nil
}

func Partition(A []int, l int, r int) int {
	p := A[l]
	i := l + 1
	for j := l + 1; j <= r; j++ {
		if A[j] < p {
			Swap(&A[j], &A[i])
			i += 1
		}
	}
	Swap(&A[l], &A[i-1])
	return i - 1
}
