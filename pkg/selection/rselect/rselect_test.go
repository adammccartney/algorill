package rselect

import "testing"

func TestRselect(t *testing.T) {
	type args struct {
		A []int
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Array with odd number elems", args{[]int{5, 16, 3, 7, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rselect(tt.args.A, tt.args.i); got != tt.want {
				t.Errorf("Rselect() = %v, want %v", got, tt.want)
			}
		})
	}
}
