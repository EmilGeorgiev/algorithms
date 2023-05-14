package math_test

import (
	"github.com/EmilGeorgiev/algorithms/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPrimeNumbersLessThen(t *testing.T) {
	cases := []struct {
		input    int64
		expected []int64
	}{
		{10, []int64{2, 3, 5, 7}},
		{20, []int64{2, 3, 5, 7, 11, 13, 17, 19}},
		{30, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{1, nil},
		{2, []int64{2}},
		{100, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, c := range cases {
		result := math.FindPrimeNumbersLessThanOrEqualTo(c.input)
		assert.Equal(t, c.expected, result)
	}
}

func BenchmarkFindPrimeNumbersLessThanOrEqualTo100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FindPrimeNumbersLessThanOrEqualTo(100)
	}
}

func BenchmarkFindPrimeNumbersLessThanOrEqualTo1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FindPrimeNumbersLessThanOrEqualTo(1000)
	}
}

func BenchmarkFindPrimeNumbersLessThanOrEqualTo10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FindPrimeNumbersLessThanOrEqualTo(10000)
		// old -   	   26083	     49068 ns/op
	}
}
