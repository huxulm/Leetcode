package sumclosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{-1, 2, 1, -4}, target: 1, expect: 2},
		{nums: []int{1, 1, 1, 0}, target: -100, expect: 2},
	}

	for _, test := range tests {
		if result := threeSumClosest(test.nums, test.target); result != test.expect {
			t.Errorf("Input nums: %v target: %d expect: %d but got: %d", test.nums, test.target, test.expect, result)
		}
	}
}
