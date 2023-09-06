package mergesort

import "testing"

func compareLists(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func TestMergeSortedLists(t *testing.T) {
	listA := []int{503, 703, 765}
	listB := []int{87, 512, 677}

	expected := []int{87, 503, 512, 677, 703, 765}

	actual, err := Merge(listA, listB)
	if err != nil {
		t.Errorf("Error merging lists: %v", err)
	}

	if !compareLists(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
