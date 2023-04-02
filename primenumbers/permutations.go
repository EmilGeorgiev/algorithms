package primenumbers

// Factorial calculates the factorial of a given non-negative integer n.
// It uses a recursive approach where Factorial(n) = n * Factorial(n-1).
// Base case: Factorial(0) = 1.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// PermutationsWithoutDuplicationCount calculates the number of permutations
// without duplication for a given set of n elements and a specified value of k.
// It uses the formula: P(n, k) = n! / (n - k)!
//
// Params:
// - n: the total number of elements in the set.
// - k: the number of elements to select from the input set for each permutation.
//
// Returns:
// - the number of unique permutations of size k that can be formed from a set of n elements.
func PermutationsWithoutDuplicationCount(n, k int) int {
	return Factorial(n) / Factorial(n-k)
}
	
// PermutationsWithoutDuplication generates all the permutations without duplication
// for a given slice of elements and a specified value of k.
//
// The function uses a recursive approach:
// 1. Base case: If k == 0, return an empty permutation.
// 2. For each element in the input slice,
//   - remove the element from the remaining elements,
//   - generate all permutations of size (k-1) using the remaining elements,
//   - append the current element to the beginning of each generated permutation,
//   - add the resulting permutation to the result set.
//
// Params:
// - elements: a slice of T representing the set of elements to permute.
// - k: the number of elements to select from the input set for each permutation.
//
// Returns:
//   - a 2D slice of T containing all permutations without duplication
//     of the input elements of size k.
func PermutationsWithoutDuplication[T any](elements []T, k int) [][]T {
	if k == 0 {
		return [][]T{{}}
	}

	var result [][]T
	for i, element := range elements {
		remainingElements := append(elements[:i], elements[i+1:]...)
		for _, perm := range PermutationsWithoutDuplication(remainingElements, k-1) {
			result = append(result, append([]T{element}, perm...))
		}
	}

	return result
}
