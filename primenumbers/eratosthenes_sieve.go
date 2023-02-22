package primenumbers

// FindPrimeNumbersLessThanOrEqualTo returns a slice of prime numbers that are
// less than or equal to the given number. The implementation uses the
// Sieve of Eratosthenes algorithm.
//
// The input number must be a positive integer greater than or equal to 2.
// If the input number is less than 2, an empty slice is returned.
//
// The function returns a slice of prime numbers.
// If no prime numbers are found, an empty slice is returned.
//
// The implementation uses an integer slice called "numbers" to keep track of
// whether each number is prime or not. If a number is marked as non-prime,
// it is skipped in subsequent iterations of the loop. When a prime number
// is found, it is added to the "result" slice. Then, all multiples of the
// prime number up to the input number are marked as non-prime in the "numbers"
// slice. This process is repeated for all remaining numbers in the "numbers"
// slice.
func FindPrimeNumbersLessThanOrEqualTo(number int64) []int64 {
	var result []int64
	numbers := make([]int64, number+1)
	for i := int64(2); i <= number; i++ {
		if numbers[i] != 0 {
			continue
		}
		result = append(result, i)
		for j := i * i; j <= number; {
			numbers[j] = 1
			j += i
		}
	}
	return result
}
