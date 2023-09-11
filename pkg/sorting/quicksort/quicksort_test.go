package quicksort

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name     string
		l        int
		r        int
		expected []int
	}{
		{"Swap 1 and 2", 1, 2, []int{2, 1}},
		{"Swap 0 and 0", 0, 0, []int{0, 0}},
		{"Swap -1 and 1", -1, 1, []int{1, -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l, r := test.l, test.r
			swap(&l, &r)
			if !reflect.DeepEqual([]int{l, r}, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, []int{l, r})
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty Slice", []int{}, []int{}},
		{"Random Order Slice",
			[]int{769, 449, 327, 513, 925, 649, 160, 713, 619, 935, 882, 392, 303, 113, 571, 890},
			[]int{113, 160, 303, 327, 392, 449, 513, 571, 619, 649, 713, 769, 882, 890, 925, 935},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Quicksort(test.input, 0, len(test.input)-1)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, test.input)
			}
		})
	}
}
