package primenumbers

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tc := range testCases {
		result := Factorial(tc.input)
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestPermutationsWithoutDuplicationCount(t *testing.T) {
	testCases := []struct {
		n, k     int
		expected int
	}{
		{5, 3, 60},
		{4, 2, 12},
		{3, 3, 6},
		{6, 0, 1},
	}

	for _, tc := range testCases {
		result := PermutationsWithoutDuplicationCount(tc.n, tc.k)
		if result != tc.expected {
			t.Errorf("PermutationsWithoutDuplicationCount(%d, %d) = %d, expected %d", tc.n, tc.k, result, tc.expected)
		}
	}
}

func TestPermutationsWithoutDuplication(t *testing.T) {
	testCases := []struct {
		elements []int
		k        int
		expected [][]int
	}{
		{
			[]int{1, 2},
			1,
			[][]int{{1}, {2}},
		},
		{
			[]int{1, 2},
			2,
			[][]int{{1, 2}, {2, 1}},
		},
		{
			[]int{1, 2, 3},
			2,
			[][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}},
		},
	}

	for _, tc := range testCases {
		result := PermutationsWithoutDuplication(tc.elements, tc.k)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("PermutationsWithoutDuplication(%v, %d) = %v, expected %v", tc.elements, tc.k, result, tc.expected)
		}
	}
}

func TestPermutationsWithDuplicationInt(t *testing.T) {
	testCases := []struct {
		elements []int
		k        int
		expected [][]int
	}{
		{
			[]int{1, 2},
			1,
			[][]int{{1}, {2}},
		},
		{
			[]int{1, 2},
			2,
			[][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
		},
		{
			[]int{1, 2, 3},
			2,
			[][]int{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}},
		},
	}

	for _, tc := range testCases {
		result := PermutationsWithDuplication(tc.elements, tc.k)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("PermutationsWithDuplicationInt(%v, %d) = %v, expected %v", tc.elements, tc.k, result, tc.expected)
		}
	}
}
