package sum

import (
	"reflect"
	"testing"
)

func TestTwoSumTarget(t *testing.T) {
	input := []int{1, 3, 1, 2, 2, 3}
	expect := [][]int{{1, 3}, {2, 2}}
	target := 4
	if r := twoSumTarget(input, target); !reflect.DeepEqual(r, expect) {
		t.Errorf("Expect: %v but got: %v", expect, r)
	}
}
