package primenumbers_test

import (
	"testing"

	"github.com/EmilGeorgiev/algorithms/primenumbers"
)

// Test for CalculatePerfectNumberByMersenne
func TestCalculatePerfectNumberByMersenne(t *testing.T) {
	testCases := []struct {
		p        int64
		expected string
	}{
		{2, "6"},
		{3, "28"},
		{5, "496"},
		{7, "8128"},
	}

	for _, testCase := range testCases {
		result, err := primenumbers.CalculatePerfectNumberByMersenne(testCase.p)
		if err != nil {
			t.Errorf("CalculatePerfectNumberByMersenne(%d) returned an error: %v", testCase.p, err)
		} else if result.String() != testCase.expected {
			t.Errorf("CalculatePerfectNumberByMersenne(%d) = %s, expected %s", testCase.p, result, testCase.expected)
		}
	}
}
