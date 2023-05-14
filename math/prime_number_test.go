package math_test

import (
	"math/big"
	"testing"

	"github.com/EmilGeorgiev/algorithms/math"
	"github.com/stretchr/testify/assert"
)

func TestIsPrimeNumber(t *testing.T) {
	// SetUp
	cases := []struct {
		name     string
		num      uint64
		expected bool
	}{
		{name: "Number 1", num: 1, expected: false},
		{name: "Number 2", num: 2, expected: true},
		{name: "Number 3", num: 3, expected: true},
		{name: "Number 4", num: 4, expected: false},
		{name: "Number 5", num: 5, expected: true},
		{name: "Number 6", num: 6, expected: false},
		{name: "Number 7", num: 7, expected: true},
		{name: "Number 8", num: 8, expected: false},
		{name: "Number 9", num: 9, expected: false},
		{name: "Number 10", num: 10, expected: false},
		{name: "Number 127", num: 127, expected: true},
		{name: "Number 8191", num: 8191, expected: true},
		{name: "Number 113", num: 113, expected: true},
		{name: "Number 140", num: 140, expected: false},
		{name: "Number 150", num: 150, expected: false},
		{name: "Number 27", num: 27, expected: false},
		{name: "Number 11", num: 11, expected: true},
		{name: "Number 15", num: 15, expected: false},
		{name: "Number 19", num: 19, expected: true},
		{name: "Number 22", num: 22, expected: false},
		{name: "Number 21", num: 21, expected: false},
		{name: "Number 31", num: 31, expected: true},
		{name: "Number 63", num: 63, expected: false},
		{name: "Number 613", num: 613, expected: true},
		{name: "Number 100", num: 100, expected: false},
		{name: "Number 104", num: 104, expected: false},
		{name: "Number 1000", num: 1000, expected: false},
		{name: "Number 1111", num: 1111, expected: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			actual := math.IsPrime(c.num)

			// Assert
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsPrimeBig(t *testing.T) {
	// Test some small prime numbers.
	for _, n := range []string{"2", "3", "5", "7", "11", "13", "17", "19", "23", "29"} {
		num := new(big.Int)
		num.SetString(n, 10)
		if !math.IsPrimeBig(num) {
			t.Errorf("%s should be prime, but IsPrimeBig returned false", n)
		}
	}

	// Test some small composite numbers.
	for _, n := range []string{"4", "6", "8", "9", "10", "12", "14", "15", "16", "18", "20", "21", "22", "24", "25"} {
		num := new(big.Int)
		num.SetString(n, 10)
		if math.IsPrimeBig(num) {
			t.Errorf("%s should not be prime, but IsPrimeBig returned true", n)
		}
	}

	// Test a large prime number.
	num := new(big.Int)
	num.SetString("6700417", 10)
	if !math.IsPrimeBig(num) {
		t.Errorf("6700417 should be prime, but IsPrimeBig returned false")
	}

	// Test a large composite number.
	num.SetString("6700427", 10)
	if math.IsPrimeBig(num) {
		t.Errorf("6700427 should not be prime, but IsPrimeBig returned true")
	}
}
