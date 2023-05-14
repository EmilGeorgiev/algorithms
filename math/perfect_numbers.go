package math

import (
	"fmt"
	"math/big"
)

// CalculatePerfectNumberByMersenne calculates a perfect number using Leonhard Euler theorem a Mersenne prime.
// A perfect number can be calculated by multiplying the Mersenne prime by one less than its corresponding power of 2.
// The function returns the perfect number and nil if the calculation is successful, or nil and an error if the input is not a Mersenne number.
//
// The input argument p is an int64 representing the exponent in the Mersenne number formula (2^p - 1).
//
// Example:
//
//	p = 3
//	The function will return 28 and nil, because (2^3 - 1) * (2^(3-1)) = 7 * 4 = 28 is a perfect number.
func CalculatePerfectNumberByMersenne(p int64) (*big.Int, error) {
	if !IsMersenneNumber(p) {
		return nil, fmt.Errorf("the number %d is not Mersenne number", p)
	}

	base := big.NewInt(2)
	exp := big.NewInt(p)
	mersennePrimeNum := new(big.Int).Exp(base, exp, nil)
	mersennePrimeNum.Sub(mersennePrimeNum, big.NewInt(1))

	exp.Sub(exp, big.NewInt(1))
	mul := new(big.Int).Exp(base, exp, nil)
	return new(big.Int).Mul(mul, mersennePrimeNum), nil
}
