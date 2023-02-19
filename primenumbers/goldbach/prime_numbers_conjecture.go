package goldbach

import "github.com/EmilGeorgiev/algorithms/primenumbers"

// RepresentEvenNumAsSumOfTwoPrimeNums is evidence that each even
// number greater than 2 can be represented as a  sum of two prime numbers.
func RepresentEvenNumAsSumOfTwoPrimeNums(evenNumber uint64) [][2]uint64 {
	var result [][2]uint64

	for i := evenNumber - 1; i >= 2; i-- {
		if primenumbers.IsPrime(i) {
			if primenumbers.IsPrime(evenNumber - i) {
				result = append(result, [2]uint64{evenNumber - i, i})
			}
		}
		if i <= evenNumber-i {
			return result
		}
	}

	return result
}
