package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expect     [][]int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7, expect: [][]int{
			{2, 2, 3}, {7},
		}},
		{candidates: []int{2, 3, 5}, target: 8, expect: [][]int{
			{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
		}},
		{candidates: []int{2}, target: 1, expect: [][]int{}},
	}

	for _, test := range tests {
		if result := combinationSum(test.candidates, test.target); !reflect.DeepEqual(test.expect, result) {
			t.Errorf("Expect: %v but got: %v", test.expect, result)
		}
	}

}
