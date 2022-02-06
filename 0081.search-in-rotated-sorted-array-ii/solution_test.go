package searchinrotatedsortedarrayii

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect bool
	}{
		{input: []int{2, 5, 6, 0, 0, 1, 2}, target: 0, expect: true},
		{input: []int{2, 5, 6, 0, 0, 1, 2}, target: 2, expect: true},
		{input: []int{2, 5, 6, 0, 0, 1, 2}, target: 3, expect: false},
		{input: []int{1, 0, 1, 1, 1, 1}, target: 0, expect: true},
		{input: []int{1, 0, 1, 1, 1, 1}, target: 1, expect: true},
		{input: []int{1, 0, 1, 1, 1, 1}, target: 2, expect: false},
	}

	for _, test := range tests {
		if result := search(test.input, test.target); result != test.expect {
			t.Errorf("Input %v target %d expect %v got %v", test.input, test.target, test.expect, result)
		}
	}
}
