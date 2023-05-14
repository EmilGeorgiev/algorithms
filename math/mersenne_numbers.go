package math

import "math/big"

// IsMersenneNumber checks if a given number (2^p - 1) is a Mersenne prime using the Lucas-Lehmer test.
// The function returns true if the number is a Mersenne prime, and false otherwise.
//
// The input argument p is an int64 representing the exponent in the Mersenne number formula (2^p - 1).
//
// Example 1:
//
//	p = 3
//	The function will return true, because 2^3 - 1 = 7 is a Mersenne prime.
//
// Example 2:
//
//	p = 11
//	The function will return false, because 2^11 - 1 = 2047 is not a Mersenne prime.
func IsMersenneNumber(p int64) bool {
	// Base case: if p is 2, return true since 2^2 - 1 = 3 is a Mersenne prime.
	if p == 2 {
		return true
	}

	// Check if p is a prime number. If p is not prime, (2^p - 1) cannot be a Mersenne prime.
	if !IsPrime(uint64(p)) {
		return false
	}

	// Calculate the Mersenne number (2^p - 1) and store it in the variable mod.
	mod := new(big.Int).Exp(big.NewInt(2), big.NewInt(p), new(big.Int))
	mod.Sub(mod, big.NewInt(1))

	// Initialize the Lucas-Lehmer sequence with the first element, 4 % mod.
	lucasLehmer := new(big.Int).Mod(big.NewInt(4), mod)
	sub := big.NewInt(2)

	// Iterate through the Lucas-Lehmer sequence (s = s^2 - 2) until the (p - 1)-th element.
	for i := int64(1); i < p-1; i++ {
		lucasLehmer = lucasLehmer.Mul(lucasLehmer, lucasLehmer)
		lucasLehmer = lucasLehmer.Sub(lucasLehmer, sub)
		lucasLehmer = lucasLehmer.Mod(lucasLehmer, mod)
	}

	// If the (p - 1)-th element of the Lucas-Lehmer sequence is 0, the Mersenne number is prime.
	return lucasLehmer.Cmp(big.NewInt(0)) == 0
}
