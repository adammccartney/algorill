package karatsuba_mult_test

import (
	"testing"

	"github.com/adammccartney/algorill/pkg/karatsuba_mult"
)

func TestKaratsubaMult(t *testing.T) {
	var tests = []struct {
		x, y, want int
	}{
		{134, 46, 6164},
		{56, 12, 672},
		{5678, 1234, 7006652},
		{95678433, 10203033, 976210209287289},
	}

	for _, test := range tests {
		got, err := karatsuba_mult.KaratsubaMult(test.x, test.y)
		if err != nil {
			t.Errorf("KaratsubaMult(%d, %d) error: %s", test.x, test.y, err)
		} else if got != test.want {
			t.Errorf("KaratsubaMult(%d, %d) = %d, want %d", test.x, test.y, got, test.want)
		}
	}
}
