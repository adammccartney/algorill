package karatsuba_mult_test

import (
	"math/big"
	"testing"

	"github.com/adammccartney/algorill/pkg/karatsuba_mult"
)

func TestKaratsubaMult(t *testing.T) {
	x, _ := new(big.Int).SetString("1234567890123456789012345678901234567890123456789012345678901234", 10)
	y, _ := new(big.Int).SetString("9876543210987654321098765432109876543210987654321098765432109876", 10)
	want, _ := new(big.Int).SetString("12193263113702179522618503273386678859451150739156363359236761158176366412964425393011121414472580028964404791648155158039986984", 10)

	got, err := karatsuba_mult.KaratsubaMult(x, y)
	if err != nil {
		t.Errorf("KaratsubaMult(%s, %s) error: %s", x.String(), y.String(), err)
	} else if got.Cmp(want) != 0 {
		t.Errorf("KaratsubaMult(%s, %s) = %s, want %s", x.String(), y.String(), got.String(), want.String())
	}
}
