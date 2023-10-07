package rselect

import "testing"

func TestRselect(t *testing.T) {
	type args struct {
		A []int
		l int
		r int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"Array with odd number elems", args{[]int{151, 190, 253, 139, 185, 236}, 0, 5, 2}, 185},
		{"Array with even num elems", args{[]int{266, 140, 145, 102, 255, 251}, 0, 4, 2}, 145},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rselect(tt.args.A, tt.args.l, tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("Rselect() = %v, want %v", got, tt.want)
			}
		})
	}
}
