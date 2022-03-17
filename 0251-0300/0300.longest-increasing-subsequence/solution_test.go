package longestincreasingsubsequence

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, expect: 4},
		{nums: []int{0, 1, 0, 3, 2, 3}, expect: 4},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, expect: 1},
	}

	for _, test := range tests {
		if result := lengthOfLIS1(test.nums); result != test.expect {
			t.Errorf("Input nums = %v expect %d but got %d", test.nums, test.expect, result)
		}
	}
}
