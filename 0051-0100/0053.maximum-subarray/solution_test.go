package maximumsubarray

import "testing"

func TestMaximumsubarray(t *testing.T) {
	cases := []struct {
		input  []int
		expect int
	}{
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expect: 6},
		{input: []int{1}, expect: 1},
		{input: []int{5, 4, -1, 7, 8}, expect: 23},
	}

	for _, c := range cases {
		result := maxSubArray(c.input)
		if result != c.expect {
			t.Errorf("Input: %v, expect: %d but got %d", c.input, c.expect, result)
		}
	}
}
