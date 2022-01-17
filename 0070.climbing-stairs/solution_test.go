package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{input: 1, expect: 1},
		{input: 2, expect: 2},
		{input: 3, expect: 3},
	}

	for _, c := range tests {
		if r := climbStairs1(c.input); r != c.expect {
			t.Errorf("Input: %d, expect: %d but got: %d", c.input, c.expect, r)
		}
	}
}
