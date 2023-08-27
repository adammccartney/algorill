package karatsuba_mult

import (
	"math"
	"math/big"
)

// split the integer x into two parts
func getParts(x *big.Int, n int) (*big.Int, *big.Int) {
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)

	pone := new(big.Int).Div(x, divisor)
	ptwo := new(big.Int).Mod(x, divisor)

	return pone, ptwo
}

// Karatsuba mult - implementation of a fast multiplication algorithm
// Discovered by Karatsuba in 1962. Bears relation to methods known to Gauss
// Input: two n-digit positive integers x and y
// Output: the product x * y
// Assumptions: n is a power of 2
func KaratsubaMult(x *big.Int, y *big.Int) (*big.Int, error) {
	zero := big.NewInt(0)

	if x.Cmp(zero) == 0 || y.Cmp(zero) == 0 {
		return zero, nil
	}

	xs := x.String()
	ys := y.String()

	// base case
	if len(xs) == 1 || len(ys) == 1 {
		result := new(big.Int).Mul(x, y)
		return result, nil
	} else {
		// check if we need to introduce padding
		if len(xs)%2 != 0 {
			xs = "0" + xs
		}
		if len(ys)%2 != 0 {
			ys = "0" + ys
		}
		m := math.Max(float64(len(xs)), float64(len(ys)))
		n := int(m / 1) // convert float to int
		ndiv2 := n / 2
		// get first and second halves of x
		a, b := getParts(x, ndiv2)
		// first and second halves of y
		c, d := getParts(y, ndiv2)
		ac, err := KaratsubaMult(a, c)
		if err != nil {
			return zero, err
		}
		bd, err := KaratsubaMult(b, d)
		if err != nil {
			return zero, err
		}
		// compute p := a + b and q := c + d
		p := new(big.Int).Add(a, b)
		q := new(big.Int).Add(c, d)
		pq, err := KaratsubaMult(p, q)
		if err != nil {
			return zero, err
		}
		adbc := new(big.Int).Sub(pq, new(big.Int).Add(ac, bd))
		powN := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
		powNdiv2 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(ndiv2)), nil)

		result := new(big.Int).Mul(ac, powN)
		result.Add(result, new(big.Int).Mul(adbc, powNdiv2))
		result.Add(result, bd)

		return result, nil
	}
}
