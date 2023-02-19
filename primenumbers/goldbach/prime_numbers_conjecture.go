package goldbach

import "github.com/EmilGeorgiev/algorithms/primenumbers"

// RepresentEvenNumAsSumOfTwoPrimeNums takes an even number as input and returns
// a slice of pairs of prime numbers that add up to the given even number.
func RepresentEvenNumAsSumOfTwoPrimeNums(evenNumber uint64) [][2]uint64 {
	var result [][2]uint64

	// Start with the largest possible prime number less than the even number and
	// check if the difference between the even number and that prime number is also prime.
	// If it is, add the pair of prime numbers to the result slice.
	for i := evenNumber - 1; i >= 2; i-- {
		if primenumbers.IsPrime(i) {
			if primenumbers.IsPrime(evenNumber - i) {
				result = append(result, [2]uint64{evenNumber - i, i})
			}
		}

		// Once 'i' is less than or equal to evenNumber-i, all possible pairs of prime numbers
		// have been checked, and we can return the result.
		if i <= evenNumber-i {
			return result
		}
	}

	// If no pairs of prime numbers that add up to the given even number were found,
	// an empty slice is returned.
	return result
}
