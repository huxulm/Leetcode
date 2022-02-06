package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 1, expect: 5},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, expect: 6},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 4, expect: 0},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 5, expect: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 6, expect: 2},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 7, expect: 3},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expect: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, expect: -1},
	}

	for _, test := range tests {
		if result := search(test.nums, test.target); result != test.expect {
			t.Errorf("Input %v target %d expect %v got %v", test.nums, test.target, test.expect, result)
		}
	}
}
