package math_test

import (
	"github.com/EmilGeorgiev/algorithms/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPrimeDivisors(t *testing.T) {
	testCases := []struct {
		input int64
		want  []int64
	}{
		{24, []int64{2, 2, 2, 3}},
		{27, []int64{3, 3, 3}},
		{29, []int64{29}},
		{30, []int64{2, 3, 5}},
		{1024, []int64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
	}

	for _, tc := range testCases {
		got := math.FindPrimeDivisors(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
