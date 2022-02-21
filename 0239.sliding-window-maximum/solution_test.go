package slidingwindowmaximum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoMax(t *testing.T) {
	tests := []struct {
		input, expect []int
	}{
		{input: []int{1, 2, 2, 3, 4, 5, 5, 6, 11, 8, 3}, expect: []int{8, 11}},
	}
	for _, test := range tests {
		if result := twoMax(test.input); !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Input %v expect %v but got %v", test.input, test.expect, result)
		} else {
			fmt.Println(result)
		}
	}
}

func twoMax(nums []int) []int {
	maxes := []int{-1 << 31, -1 << 31}
	for i := range nums {
		if nums[i] > maxes[1] {
			maxes[0], maxes[1] = maxes[1], nums[i]
		} else if nums[i] > maxes[0] {
			maxes[0] = nums[i]
		}
	}
	return maxes
}

func TestSlidingWindow(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		expect []int
	}{
		{input: []int{2, 3, 5, -3, 8, -2, -4, 0}, k: 2},
	}

	for _, test := range tests {
		fmt.Println(maxSlidingWindow(test.input, test.k))
		fmt.Println(minSlidingWindow(test.input, test.k))
	}
}
