package math

import (
	"math"
	"math/big"
)

var two = big.NewInt(2)

// IsPrime checks if a given number is a prime number or not.
// It returns true if the number is prime, and false otherwise.
func IsPrime(num uint64) bool {
	// Check if the number is less than 2, which is not a prime number.
	if num < 2 {
		return false
	}

	// Check if the number is 2, which is a prime number.
	if num == 2 {
		return true
	}

	// Check if the number is even, which is not a prime number.
	if num%2 == 0 {
		return false
	}

	// Check if the number is divisible by any odd number less than or equal to its square root.
	// If it is, then the number is not a prime number.
	f := uint64(math.Sqrt(float64(num)))
	for i := uint64(3); i <= f; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	// If none of the above conditions are met, then the number is a prime number.
	return true
}

// IsPrimeBig checks if a given big integer is a prime number.
func IsPrimeBig(num *big.Int) bool {
	// Check if the number is less than 2, which is not a prime number.
	if num.Cmp(big.NewInt(2)) == -1 {
		return false
	}

	// Check if the number is 2, which is a prime number.
	if num.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// Check if the number is even, which is not a prime number.
	bigI := new(big.Int).Mod(num, two)
	if bigI.Cmp(big.NewInt(0)) == 0 {
		return false
	}

	// Calculate the square root of the input number and set up a divisor to test potential factors.
	bigI.Sqrt(num)
	sqrt := big.NewInt(0)
	sqrt.Set(bigI)

	divisor := big.NewInt(3)
	for divisor.Cmp(sqrt) < 1 {
		// Check if the current divisor is a factor of the input number.
		bigI.Mod(num, divisor)
		if rr := bigI.Cmp(big.NewInt(0)); rr == 0 {
			return false
		}
		// Increment the divisor by 2 to skip even divisors and optimize the loop.
		divisor.Add(divisor, two)
	}

	// If none of the above conditions are met, then the number is a prime number.
	return true
}
