package rselect

import (
	"math/rand"

	"github.com/adammccartney/algorill/pkg/sorting/quicksort"
)

func partition(A []int, l int, r int, p int) int {
	pval := A[p]
	quicksort.Swap(&A[p], &A[r])
	storeIndex := l
	for i := l; i < r; i++ {
		if A[i] < pval {
			quicksort.Swap(&A[storeIndex], &A[i]) // move pivot to final pos
			storeIndex += 1
		}
	}
	quicksort.Swap(&A[r], &A[storeIndex])
	return storeIndex
}

func Rselect(A []int, l int, r int, k int) int {
	if l == r { // list only contains 1 elem
		return A[l]
	}
	// choose pivot p uniformly at random
	p := l + rand.Intn(r-l+1)
	j := partition(A, l, r, p)
	if k == j { // got lucky
		return A[k]
	} else if k < j {
		return Rselect(A, l, j-1, k)
	} else {
		return Rselect(A, l+1, r, k)
	}
}
