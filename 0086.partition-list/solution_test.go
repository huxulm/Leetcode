package partitionlist

import (
	"reflect"
	"testing"

	. "lc/structures"
	"lc/testutils"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		input  *ListNode
		x      int
		expect *ListNode
	}{
		{input: testutils.Arr2List([]int{1, 4, 3, 2, 5, 2}), x: 3, expect: testutils.Arr2List([]int{1, 2, 2, 4, 3, 5})},
	}

	for _, test := range tests {
		if result := partition(test.input, test.x); !reflect.DeepEqual(testutils.List2Arr(test.expect), testutils.List2Arr(result)) {
			t.Errorf("Input x %d expect %v but got %v", test.x, testutils.List2Arr(test.expect), testutils.List2Arr(result))
		}
	}
}
