package eratosthenes_test

import (
	"testing"

	"github.com/EmilGeorgiev/algorithms/primenumbers/eratosthenes"
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
		result := eratosthenes.FindPrimeNumbersUpTo(c.input)
		assert.Equal(t, c.expected, result)
	}
}
