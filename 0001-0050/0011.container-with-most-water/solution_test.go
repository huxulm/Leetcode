package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expect: 49},
	}

	for _, c := range tests {
		if r := maxArea(c.input); r != c.expect {
			t.Errorf("Input: %v expect: %d but got: %d", c.input, c.expect, r)
		}
	}
}
