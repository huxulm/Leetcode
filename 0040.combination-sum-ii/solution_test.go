package combinationsumii

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect [][]int
	}{
		{input: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, expect: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{input: []int{2, 5, 2, 1, 2}, target: 5, expect: [][]int{{1, 2, 2}, {5}}},
	}

	for _, test := range tests {
		if result := combinationSum2(test.input, test.target); !reflect.DeepEqual(test.expect, result) {
			t.Errorf("Input: %v expect %v but got %v", test.input, test.expect, result)
		}
	}
}
