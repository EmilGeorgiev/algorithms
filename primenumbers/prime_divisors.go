package primenumbers

// FindPrimeDivisors returns the prime factors of the input integer 'n'.
func FindPrimeDivisors(n int64) []int64 {
	var result []int64
	i := int64(1)

	// Keep looping until the number is fully factored.
	for n != 1 {
		// Advance to the next prime number.
		i++

		// If the current prime number divides the remaining factor,
		// append it to the result slice and divide the factor by the prime.
		for n%i == 0 {
			result = append(result, i)
			n /= i
		}
	}

	return result
}
