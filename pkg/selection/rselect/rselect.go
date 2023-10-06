package rselect

import (
	"log"

	"github.com/adammccartney/algorill/pkg/sorting/quicksort"
)

func Rselect(A []int, i int) int {
	if len(A) == 1 {
		return A[0]
	}
	// choose pivot p uniformly at random
	p, err := quicksort.ChoosePivot(A, 0, len(A)-1)
	if err != nil {
		log.Fatal(err)
	}
	j := quicksort.Partition(A, p, len(A)-1)
	if j == i { // got lucky
		return A[p]
	} else if j > i {
		return Rselect(A[:p], i)
	} else {
		return Rselect(A[p+1:], i-j)
	}
}
