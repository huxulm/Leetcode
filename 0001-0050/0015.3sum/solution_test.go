package sum

import (
	"reflect"
	"testing"
)

func TestTwoSumTarget(t *testing.T) {
	input := []int{1, 3, 1, 2, 2, 3}
	expect := [][]int{{1, 3}, {2, 2}}
	target := 4
	if r := twoSumTarget(input, 0, target); !reflect.DeepEqual(r, expect) {
		t.Errorf("Expect: %v but got: %v", expect, r)
	}
}
func TestThreeSumTarget(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	expect := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if r := threeSum(input); !reflect.DeepEqual(r, expect) {
		t.Errorf("Expect: %v but got: %v", expect, r)
	}
}
