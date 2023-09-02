package countinv_test

import (
	"reflect"
	"testing"

	"github.com/adammccartney/algorill/pkg/countinv"
)

type result struct {
	output []int
	inv    int
}

// Checks that the number of inversions in the input array is correct.
func TestCountInv(t *testing.T) {
	tests := []struct {
		input    []int
		output   []int
		expected int
	}{
		{[]int{1, 3, 2, 4}, []int{1, 2, 3, 4}, 1},
		{[]int{1, 5, 4, 2}, []int{1, 2, 4, 5}, 3},
		{[]int{5, 1, 3, 2, 4, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 5},
	}

	for _, test := range tests {
		var res result
		arr, inv, _ := countinv.SortCountInv(test.input)
		res.output = arr
		res.inv = inv
		if !reflect.DeepEqual(test.output, res.output) {
			t.Errorf("Expected %v, got %v", test.output, res.output)
		}
		if test.expected != res.inv {
			t.Errorf("Expected %v, got %v", test.expected, res.inv)
		}
	}
}
