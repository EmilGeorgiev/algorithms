package math

func FindCombinations[T any](elements []T, k int64) [][]T {

	n := len(elements)

	data := make([]T, k)

	//result := combinations(arr, data, 0, n-1, 0, k)
	return findCombinations(elements, data, 0, int64(n-1), 0, k)
}

func findCombinations[T any](elements []T, data []T, start, end, index, k int64) (result [][]T) {
	if index == k {
		combination := make([]T, k)
		copy(combination, data)
		return [][]T{combination}
	}

	for i := start; i <= end && end-i+1 >= k-index; i++ {
		data[index] = elements[i]
		combs := findCombinations(elements, data, i+1, end, index+1, k)
		result = append(result, combs...)
	}
	return
}
