package primenumbers

import "math"

// IsPrime check whether a number is a prime number.
//
// Always when 'num' has divider greater than square root of 'num', this means that
// 'num' has divider lower than square root of num.
func IsPrime(num uint64) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	f := uint64(math.Sqrt(float64(num)))
	for i := uint64(3); i <= f; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
