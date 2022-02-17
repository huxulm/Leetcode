package sum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect [][]int
	}{
		{
			nums:   []int{-5, -3, 0, 1, 1, 2, 5, 6, 7},
			expect: [][]int{{-5, -3, 1, 7}, {-5, -3, 2, 6}, {-3, 0, 1, 2}},
		},
	}

	for _, test := range tests {
		if result := fourSum(test.nums, 0); !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Input %v expect %v but got %v", test.nums, test.expect, result)
		}
	}
}
