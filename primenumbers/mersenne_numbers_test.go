package primenumbers

import "testing"

func TestIsMersenne(t *testing.T) {
	testCases := []struct {
		p        int64
		expected bool
	}{
		{2, true}, // 2^2 - 1 = 3 is prime
		{3, true}, // 2^3 - 1 = 7 is prime
		{4, false},
		{5, true}, // 2^5 - 1 = 31 is prime
		{6, false},
		{7, true}, // 2^7 - 1 = 127 is prime
		{8, false},
		{9, false},
		{10, false},
		{11, false}, // 2^11 - 1 = 2047 is composite
		{13, true},  // 2^13 - 1 = 8191 is prime
		{15, false},
		{17, true}, // 2^17 - 1 = 131071 is prime
		{19, true}, // 2^19 - 1 = 524287 is prime, but not a Mersenne number
		{23, false},
		{29, false}, // 2^29 - 1 = 536870911 is prime, but not a Mersenne number
		{31, true},  // 2^31 - 1 = 2147483647 is prime
		{37, false},
		{43, false},
		{53, false},
		{61, true},
		{89, true},
		{107, true},
		{127, true},
		{521, true},
		{607, true},
		{1279, true},
		{2203, true},
		{2281, true},
		{3217, true},
		{4253, true},
		{4423, true},
		{9689, true},
		{9941, true},
		{11213, true},
	}

	for _, tc := range testCases {
		result := IsMersenneNumber(tc.p)
		if result != tc.expected {
			t.Errorf("IsMersenne(%d) = %v; want %v", tc.p, result, tc.expected)
		}
	}
}
