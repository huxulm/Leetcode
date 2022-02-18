package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expect: 6},
	}

	for _, test := range tests {
		if result := maxSubArray(test.nums); test.expect != result {
			t.Errorf("Input %v expect %d but got %d", test.nums, test.expect, result)
		}
	}
}
