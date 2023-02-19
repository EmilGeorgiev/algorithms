package goldbach

import "github.com/EmilGeorgiev/algorithms/primenumbers"

// RepresentEvenNumAsSumOfTwoPrimeNums is evidence that each even
// number greater than 2 can be represented as a  sum of two prime numbers.
func RepresentEvenNumAsSumOfTwoPrimeNums(evenNumber uint64) [][2]uint64 {
	var primes []uint64
	for i := uint64(2); i < evenNumber; i++ {
		if primenumbers.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	var result [][2]uint64
	for i := 0; i < len(primes); i++ {
		for j := i; j < len(primes); j++ {
			if primes[i]+primes[j] == evenNumber {
				result = append(result, [2]uint64{primes[i], primes[j]})
			}
		}
	}

	return result
}
