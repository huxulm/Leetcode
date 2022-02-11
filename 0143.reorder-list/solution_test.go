package reorderlist

import (
	"reflect"
	"testing"

	. "lc/structures"
	"lc/testutils"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input  *ListNode
		expect []int
	}{
		{input: testutils.Arr2List([]int{1, 2, 3, 4}), expect: []int{1, 4, 2, 3}},
	}

	for _, test := range tests {
		reorderList(test.input)
		if !reflect.DeepEqual(testutils.List2Arr(test.input), test.expect) {
			t.Errorf("Expect %v got %v", test.expect, testutils.List2Arr(test.input))
		}
	}
}
