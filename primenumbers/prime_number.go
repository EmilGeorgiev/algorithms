package primenumbers

import "math"

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
