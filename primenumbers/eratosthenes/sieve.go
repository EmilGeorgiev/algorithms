package eratosthenes

// FindPrimeNumbersLessThanOrEqualTo return all prime number less than parameter
// 'number'. The functions use the Eratosthenes's sieve algorithm.
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
