package jumpgame

import "testing"

func TestJumpGame(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{input: []int{0}, expect: true},
		{input: []int{2, 3, 1, 1, 4}, expect: true},
		{input: []int{2, 3, 0, 1, 4}, expect: true},
		{input: []int{2, 3, 1, 1, 0}, expect: true},
		{input: []int{2, 5, 0, 0}, expect: true},
		{input: []int{3, 0, 8, 2, 0, 0, 1}, expect: true},
		{input: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, expect: true},
	}

	for _, test := range tests {
		if result := canJump(test.input); result != test.expect {
			t.Errorf("Input %v expect %v but got %v", test.input, test.expect, result)
		}
	}
}
