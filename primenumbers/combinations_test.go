package primenumbers_test

import (
	"github.com/EmilGeorgiev/algorithms/primenumbers"
	"reflect"
	"testing"
)

func TestGetCombinations(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		k        int64
		expected [][]int
	}{
		{
			name:     "Test C(5, 4)",
			elements: []int{1, 2, 3, 4, 5},
			k:        4,
			expected: [][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 5},
				{1, 2, 4, 5},
				{1, 3, 4, 5},
				{2, 3, 4, 5},
			},
		},
		{
			name:     "Test C(3, 2)",
			elements: []int{1, 2, 3},
			k:        2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := primenumbers.FindCombinations(tt.elements, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
