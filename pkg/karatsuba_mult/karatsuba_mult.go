package karatsuba_mult

import (
	"math"
	"strconv"
)

// split the integer x into two parts
func getParts(x, n int) (int, int) {
	divisor := int(math.Pow10(n))

	if x >= divisor {
		pone := x / divisor
		ptwo := x % divisor
		return pone, ptwo
	} else {
		return 0, x
	}
}

// Karatsuba mult - implementation of a fast multiplication algorithm
// Discovered by Karatsuba in 1962. Bears relation to methods known to Gauss
// Input: two n-digit positive integers x and y
// Output: the product x * y
// Assumptions: n is a power of 2
func KaratsubaMult(x int, y int) (int, error) {

	if x == 0 || y == 0 {
		return 0, nil
	}

	xs := strconv.Itoa(x)
	ys := strconv.Itoa(y)
	// base case
	if len(xs) == 1 || len(ys) == 1 {
		return x * y, nil
	} else {
		// setup pivot points for number strings
		// prepend zero if uneven number of digits
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
			return 0, err
		}
		bd, err := KaratsubaMult(b, d)
		if err != nil {
			return 0, err
		}
		// compute p := a + b and q := c + d
		p := a + b
		q := c + d
		pq, err := KaratsubaMult(p, q)
		if err != nil {
			return 0, err
		}
		adbc := pq - ac - bd
		result := ac*int(math.Pow10(n)) + adbc*int(math.Pow10(ndiv2)) + bd
		return result, nil
	}
}
