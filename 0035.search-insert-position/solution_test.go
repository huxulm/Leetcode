package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		expect int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			expect: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			expect: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			expect: 4,
		},
		{
			nums:   []int{1},
			target: 0,
			expect: 0,
		},
	}

	for _, c := range cases {
		result := searchInsert(c.nums, c.target)
		if result != c.expect {
			t.Errorf("Input: %v, target: %d expect %d but got: %d", c.nums, c.target, c.expect, result)
		}
	}
}
