package primenumbers_test

import (
	"github.com/EmilGeorgiev/algorithms/primenumbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepresentEvenNumAsSumOfTwoPrimeNums(t *testing.T) {
	// SetUp
	cases := []struct {
		name       string
		evenNumber uint64
		expected   [][2]uint64
	}{
		{name: "Even number 4", evenNumber: 4, expected: [][2]uint64{{2, 2}}},
		{name: "Even number 6", evenNumber: 6, expected: [][2]uint64{{3, 3}}},
		{name: "Even number 8", evenNumber: 8, expected: [][2]uint64{{3, 5}}},
		{name: "Even number 10", evenNumber: 10, expected: [][2]uint64{{3, 7}, {5, 5}}},
		{name: "Even number 12", evenNumber: 12, expected: [][2]uint64{{5, 7}}},
		{name: "Even number 14", evenNumber: 14, expected: [][2]uint64{{3, 11}, {7, 7}}},
		{name: "Even number 16", evenNumber: 16, expected: [][2]uint64{{3, 13}, {5, 11}}},
		{name: "Even number 18", evenNumber: 18, expected: [][2]uint64{{5, 13}, {7, 11}}},
		{name: "Even number 28", evenNumber: 28, expected: [][2]uint64{{5, 23}, {11, 17}}},
		{name: "Even number 100", evenNumber: 100, expected: [][2]uint64{{3, 97}, {11, 89}, {17, 83}, {29, 71}, {41, 59}, {47, 53}}},
	}
	// 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			actual := primenumbers.RepresentEvenNumAsSumOfTwoPrimeNums(c.evenNumber)

			// Assert
			assert.Equal(t, c.expected, actual)
		})
	}
}
