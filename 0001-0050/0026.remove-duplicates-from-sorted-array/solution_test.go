package removeduplicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		inputs []int
		expect int
	}{
		{inputs: []int{1, 1}, expect: 1},
		{inputs: []int{1, 2, 2}, expect: 2},
		{inputs: []int{1, 2, 3, 3, 3, 4, 4, 5, 5}, expect: 5},
	}

	for _, test := range tests {
		if result := removeDuplicates(test.inputs); test.expect != result {
			t.Errorf("Input %v expect %d but got %d", test.inputs, test.expect, result)
		}
	}
}
