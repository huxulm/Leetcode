package sqrtx

import "testing"

func TestSqrtx(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{8, 2},
		{9, 3},
		{12, 3},
		{12, 3},
		{17, 4},
		{25, 5},
		{31, 5},
	}

	for _, c := range tests {
		if r := mySqrt(c.input); c.expect != r {
			t.Errorf("Sqrt(%d) expect %d but got: %d", c.input, c.expect, r)
		}
	}
}
